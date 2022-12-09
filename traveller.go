package traveller

import "reflect"

// Represents a single traversal.
type Traversal struct {
	rv       reflect.Value
	parentRv reflect.Value
	key      any
	next     func(reflect.Value) (keepSearching bool)
}

// Get the currently traversed value.
//
// The value is not "unboxed" and will need to be inspected manually.
func (t Traversal) RV() reflect.Value {
	return t.rv
}

// Get the parent of the current value where this value originated.
//
// The parent value is usually "unboxed" and will represent the direct type.
func (t Traversal) ParentRV() reflect.Value {
	return t.parentRv
}

// Get the key that is used to obtain the value from the parent.
//
// The type of key depends on parentRv's kind:
// reflect.Map is reflect.Value, reflect.Struct is string, reflect.Array is int.
func (t Traversal) Key() any {
	return t.key
}

// Continue to the next traversal. Returns true if traversal should continue.
func (t Traversal) Next(rv reflect.Value) bool {
	return t.next(rv)
}

// The callback on each traversal.
//
// Return true on the second return value to continue traversal.
type TraversalFunc func(Traversal) (keepSearching bool)

// Represents a single finding.
type Found struct {
	// The current value.
	rv reflect.Value

	// The parent of the current value where this value originated.
	parentRv reflect.Value

	// The type of key depends on parentRv's kind:
	// reflect.Map is reflect.Value, reflect.Struct is string, reflect.Array is int.
	key any
}

// Get the currently traversed value.
//
// The value is not "unboxed" and will need to be inspected manually.
func (f Found) RV() reflect.Value {
	return f.rv
}

// Get the parent of the current value where this value originated.
//
// The parent value is usually "unboxed" and will represent the direct type.
func (f Found) ParentRV() reflect.Value {
	return f.parentRv
}

// Get the key that is used to obtain the value from the parent.
//
// The type of key depends on parentRv's kind:
// reflect.Map is reflect.Value, reflect.Struct is string, reflect.Array is int.
func (f Found) Key() any {
	return f.key
}

// The callback on each found value.
//
// Return true to continue traversal.
type FoundFunc func(Found) (keepSearching bool)

// The traveller that is used to coordinate traversal through a value.
type Traveller struct {
	// The matcher path.
	mp []Matcher

	// Callbacks for the traveller.
	cb TravellerCallback

	noFlatEmbeds bool
	ignoreStruct bool
	ignoreMap    bool
	ignoreArray  bool
}

// The list of callbacks that the traveller can call on specific events.
type TravellerCallback struct {
	// The handler to trigger on each traversal.
	OnTraversal TraversalFunc

	// The handler to trigger when a matching value is found.
	OnFound FoundFunc
}

// Manually start a new traversal using the given value, path, and callbacks.
func StartTraversal(rv reflect.Value, mp []Matcher, cb TravellerCallback, options ...TravellerOption) {
	traveller := &Traveller{
		mp: mp,
		cb: cb,
	}
	traveller.applyOptions(options)

	traveller.Match(0, rv, reflect.Value{}, nil)
}

// Applies the list of options to the traveller.
func (t *Traveller) applyOptions(options []TravellerOption) {
	for _, option := range options {
		option(t)
	}
}

// Match at a specific path element with the given value.
func (t *Traveller) Match(index int, rv, parentRv reflect.Value, key any) (keepSearching bool) {
	next := func(newRv reflect.Value) bool {
		if index == len(t.mp) {
			return t.cb.OnFound == nil || t.cb.OnFound(Found{rv: newRv, parentRv: parentRv, key: key})
		}

		segment := MatcherSegment{
			traveller: t,
			index:     index,
		}
		return t.mp[index].Match(newRv, segment)
	}

	if t.cb.OnTraversal != nil {
		return t.cb.OnTraversal(Traversal{
			rv:       rv,
			parentRv: parentRv,
			key:      key,
			next:     next,
		})
	}

	return next(rv)
}

// Whether to not flatten embedded values in structs when matching.
func (t Traveller) NoFlatEmbeds() bool {
	return t.noFlatEmbeds
}

// Whether to ignore structs on traversal.
func (t Traveller) IgnoreStruct() bool {
	return t.ignoreStruct
}

// Whether to ignore maps on traversal.
func (t Traveller) IgnoreMap() bool {
	return t.ignoreMap
}

// Whether to ignore arrays on traversal.
func (t Traveller) IgnoreArray() bool {
	return t.ignoreArray
}
