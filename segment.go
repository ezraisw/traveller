package traveller

import "reflect"

type MatcherSegment struct {
	traveller *Traveller
	index     int
}

// Get the traveller instance.
func (s MatcherSegment) Traveller() *Traveller {
	return s.traveller
}

// Get the current index of the path.
func (s MatcherSegment) Index() int {
	return s.index
}

// Go to the next path segment with a new value.
//
// False is returned when traversal should not be continued.
func (s MatcherSegment) Next(rv reflect.Value, parentRv reflect.Value, key any) bool {
	return s.Traveller().Match(s.Index()+1, rv, parentRv, key)
}

// Stay on the current path segment while inspecting a new value.
//
// False is returned when traversal should not be continued.
func (s MatcherSegment) Stay(rv reflect.Value, parentRv reflect.Value, key any) bool {
	return s.Traveller().Match(s.Index(), rv, parentRv, key)
}
