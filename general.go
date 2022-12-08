package traveller

import (
	"reflect"
)

const (
	PanicMsgNoMatch           = "traveller: no match"
	PanicMsgNotAPointerForSet = "traveller: not a pointer, cannot set without pointer value"
)

// The callback for setting value.
//
// Return the shouldSet as false to not set.
type SetterFunc[T any] func(oldVal T) (newVal any, shouldSet bool)

// Get first value of type T, matching path.
func MustGet[T any](i any, mp []Matcher, options ...TravellerOption) T {
	val, ok := Get[T](i, mp)
	if !ok {
		panic(PanicMsgNoMatch)
	}
	return val
}

// Get first value of type T, matching path.
// The second value will be false if nothing is found.
func Get[T any](i any, mp []Matcher, options ...TravellerOption) (T, bool) {
	var (
		val T
		ok  bool
	)

	onFound := func(f Found) bool {
		if val, ok = f.RV().Interface().(T); ok {
			return false // Stop searching on first match.
		}
		return true // Keep searching.
	}

	StartTraversal(reflect.ValueOf(i), mp, TravellerCallback{OnFound: onFound}, options...)
	return val, ok
}

// Get all value of type T, matching path.
func GetAll[T any](i any, mp []Matcher, options ...TravellerOption) []T {
	vals := make([]T, 0)

	onFound := func(f Found) bool {
		vals = appendOnTypeMatch(vals, f.RV())
		return true // Keep searching.
	}

	StartTraversal(reflect.ValueOf(i), mp, TravellerCallback{OnFound: onFound}, options...)
	return vals
}

func Set(in any, mp []Matcher, val any, options ...TravellerOption) bool {
	return SetBy(in, mp, func(any) (any, bool) { return val, true }, options...)
}

func SetBy[T any](in any, mp []Matcher, setter SetterFunc[T], options ...TravellerOption) bool {
	inRv := reflect.ValueOf(in)
	if inRv.Kind() != reflect.Ptr {
		panic(PanicMsgNotAPointerForSet)
	}
	inRv = inRv.Elem()

	changed := false
	cb := TravellerCallback{
		OnTraversal: handleInaddrVals,
		OnFound: func(f Found) bool {
			oldVal, ok := f.RV().Interface().(T)
			if !ok {
				return true // Keep searching.
			}

			newVal, shouldSet := setter(oldVal)
			if !shouldSet {
				return true // Keep searching.
			}

			newRv := reflect.ValueOf(newVal)

			// Only set compatible types.
			if rv := f.RV(); newRv.Type().AssignableTo(rv.Type()) {
				rv.Set(newRv)
				changed = true
			}

			return !changed
		},
	}

	StartTraversal(inRv, mp, cb, options...)
	return changed
}

func SetAll(in any, mp []Matcher, val any, options ...TravellerOption) int {
	return SetAllBy(in, mp, func(any) (any, bool) { return val, true }, options...)
}

func SetAllBy[T any](in any, mp []Matcher, setter SetterFunc[T], options ...TravellerOption) int {
	inRv := reflect.ValueOf(in)
	if inRv.Kind() != reflect.Ptr {
		panic(PanicMsgNotAPointerForSet)
	}
	inRv = inRv.Elem()

	count := 0
	cb := TravellerCallback{
		OnTraversal: handleInaddrVals,
		OnFound: func(f Found) bool {
			oldVal, ok := f.RV().Interface().(T)
			if !ok {
				return true // Keep searching.
			}

			newVal, shouldSet := setter(oldVal)
			if !shouldSet {
				return true // Keep searching.
			}

			newRv := reflect.ValueOf(newVal)

			// Only set compatible types.
			if rv := f.RV(); newRv.Type().AssignableTo(rv.Type()) {
				rv.Set(newRv)
				count++
			}

			return true // Keep searching.
		},
	}

	StartTraversal(inRv, mp, cb, options...)
	return count
}

func appendOnTypeMatch[T any](slice []T, rv reflect.Value) []T {
	if v, ok := rv.Interface().(T); ok {
		slice = append(slice, v)
	}
	return slice
}

// Whether the given value is a mutable but inaddressable type contained in an interface.
func stackMutable(kind reflect.Kind) bool {
	return kind == reflect.Struct || kind == reflect.Array
}

// Whether the given value is a mutable but inaddressable type contained in an interface.
func mutableInaddr(rv reflect.Value) bool {
	return rv.Kind() == reflect.Interface && stackMutable(rv.Elem().Kind())
}

func handleInaddrVals(t Traversal) bool {
	if t.RV().CanAddr() && !mutableInaddr(t.RV()) {
		return t.Next(t.RV())
	}

	// Workaround for things that return inaddressable values.
	// Copy the value as a new value and substitute it for next traversals.
	// Assign it back as the new value to the parent.
	// TODO: Find a way to prevent substitution of the value WITHOUT reflect.DeepEqual if we don't need to.

	newRv := reflect.New(t.RV().Type()).Elem()

	// Handle these inaddressable but editable types.
	// This is a workaround to enable edits of non pointer struct and array behind an interface.
	// Ideally, traversal of these values should be passed though Unbox first.
	var hackRv reflect.Value
	if mutableInaddr(t.RV()) {
		// Pointer of the value inside the interface.
		hackRv = reflect.New(t.RV().Elem().Type())
		hackRv.Elem().Set(t.RV().Elem())

		newRv.Set(hackRv)
	} else {
		newRv.Set(t.RV())
	}

	keepSearching := t.Next(newRv)

	// Restore the hack so that the types are the same.
	// Only when mutableInaddr is true.
	if hackRv.IsValid() {
		newRv.Set(hackRv.Elem())
	}

	switch parentRv := t.ParentRV(); parentRv.Kind() {
	case reflect.Struct:
		fieldName := t.Key().(string)
		parentRv.FieldByName(fieldName).Set(newRv)
	case reflect.Map:
		keyRv := t.Key().(reflect.Value)
		parentRv.SetMapIndex(keyRv, newRv)
	case reflect.Array, reflect.Slice:
		i := t.Key().(int)
		parentRv.Index(i).Set(newRv)
	}

	return keepSearching
}
