package traveller

import (
	"reflect"
)

const (
	panicMsgNoMatch           = "traveller: no match"
	panicMsgNotAPointerForSet = "traveller: not a pointer, cannot set without pointer value"
)

// The callback for setting value.
// The type returned could be anything, as long as it
// is assignable to the field in context.
//
// Return `keepSearching` as false to stop traversing.
// Return `shouldSet` as false to not set the current matched value.
type SetterFunc[T any] func(oldVal T) (newVal any, keepSearching, shouldSet bool)

// Get first value of type T, matching path.
//
// Will panic if there is no match.
func MustGet[T any](i any, mp []Matcher, options ...TravellerOption) T {
	val, ok := Get[T](i, mp)
	if !ok {
		panic(panicMsgNoMatch)
	}
	return val
}

// Get first value of type T, matching path.
//
// The second value will be false if there is no match of path and type.
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

// Set a single value matching the path using the given value.
// It will only assign once. If unsuccessful in setting the value
// on a matching field, it will continue to the next matching field.
//
// `in` must be a pointer to a value or it will panic.
func Set(in any, mp []Matcher, val any, options ...TravellerOption) bool {
	return SetBy(in, mp, func(any) (any, bool, bool) { return val, true, true }, options...)
}

// Set a single value using a function matching the path and type.
//
// Return true as the second return value for the `setter` to continue
// searching.
// Return true as the third return value for the `setter` to attempt
// a value assignment and cause the traversal to stop if
// the value is successfully set.
//
// `in` must be a pointer to a value or it will panic.
func SetBy[T any](in any, mp []Matcher, setter SetterFunc[T], options ...TravellerOption) bool {
	inRv := reflect.ValueOf(in)
	if inRv.Kind() != reflect.Ptr {
		panic(panicMsgNotAPointerForSet)
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

			newVal, keepSearching, shouldSet := setter(oldVal)
			if !shouldSet {
				return keepSearching // Keep searching.
			}

			newRv := reflect.ValueOf(newVal)

			// Only set compatible types.
			if rv := f.RV(); newRv.Type().AssignableTo(rv.Type()) {
				rv.Set(newRv)
				changed = true
			}

			return keepSearching && !changed
		},
	}

	StartTraversal(inRv, mp, cb, options...)
	return changed
}

// Set all fields matching the path using the given value.
// Will only assign the value if it is assignable to the matching field.
//
// `in` must be a pointer to a value or it will panic.
func SetAll(in any, mp []Matcher, val any, options ...TravellerOption) int {
	return SetAllBy(in, mp, func(any) (any, bool, bool) { return val, true, true }, options...)
}

// Set all values using a function matching the path and type.
//
// Return true as the second return value for the `setter` to
// continue searching.
// Return true as the third return value for the `setter` to set
// the desired value.
//
// `in` must be a pointer to a value or it will panic.
func SetAllBy[T any](in any, mp []Matcher, setter SetterFunc[T], options ...TravellerOption) int {
	inRv := reflect.ValueOf(in)
	if inRv.Kind() != reflect.Ptr {
		panic(panicMsgNotAPointerForSet)
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

			newVal, keepSearching, shouldSet := setter(oldVal)
			if !shouldSet {
				return keepSearching
			}

			newRv := reflect.ValueOf(newVal)

			// Only set compatible types.
			if rv := f.RV(); newRv.Type().AssignableTo(rv.Type()) {
				rv.Set(newRv)
				count++
			}

			return keepSearching
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

// Whether it is a mutable type but is stack allocated by default.
// This characteristic is applied to arrays and structs.
func stackMutable(kind reflect.Kind) bool {
	return kind == reflect.Struct || kind == reflect.Array
}

// Whether the given value is a mutable but inaddressable type contained in an interface.
func mutableInaddr(rv reflect.Value) bool {
	return rv.Kind() == reflect.Interface && stackMutable(rv.Elem().Kind())
}

// The handler for handling nested inaddressable values.
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

	// The last condition is to fix a problem when assigning a value to an any-typed field
	// that currently has a struct or an array value. If the index is the last, then it is
	// most possibly the one we're trying to assign into instead of modifying the members.
	if mutableInaddr(t.RV()) && t.Index() != t.Traveller().PathLen() {
		// Pointer of the value inside the interface.
		hackRv = reflect.New(t.RV().Elem().Type())
		hackRv.Elem().Set(t.RV().Elem())

		newRv.Set(hackRv)
	} else {
		newRv.Set(t.RV())
	}

	keepSearching := t.Next(newRv)

	// Restore the hack so that the types are the same.
	// Only when mutableInaddr() returns true and is not the found value.
	if hackRv.IsValid() {
		newRv.Set(hackRv.Elem())
	}

	setForParent(t.ParentRV(), t.Key(), newRv)

	return keepSearching
}

func setForParent(parentRv reflect.Value, key any, newRv reflect.Value) {
	switch parentRv.Kind() {
	case reflect.Struct:
		fieldName := key.(string)
		parentRv.FieldByName(fieldName).Set(newRv)
	case reflect.Map:
		keyRv := key.(reflect.Value)
		parentRv.SetMapIndex(keyRv, newRv)
	case reflect.Array, reflect.Slice:
		i := key.(int)
		parentRv.Index(i).Set(newRv)
	}
}
