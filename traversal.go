package traveller

import "reflect"

// Represents a single traversal.
type Traversal struct {
	traveller *Traveller
	index     int
	rv        reflect.Value
	parentRv  reflect.Value
	key       any
	next      func(reflect.Value) (keepSearching bool)
}

// Get the traveller instance.
func (t Traversal) Traveller() *Traveller {
	return t.traveller
}

// Get the current index of the path.
func (t Traversal) Index() int {
	return t.index
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
