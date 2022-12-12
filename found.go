package traveller

import "reflect"

// Represents a single finding.
type Found struct {
	traveller *Traveller
	rv        reflect.Value
	parentRv  reflect.Value
	key       any
}

// Get the traveller instance.
func (t Found) Traveller() *Traveller {
	return t.traveller
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
