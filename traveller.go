package traveller

import "reflect"

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
			return t.cb.OnFound == nil || t.cb.OnFound(Found{traveller: t, rv: newRv, parentRv: parentRv, key: key})
		}

		segment := MatcherSegment{
			traveller: t,
			index:     index,
		}
		return t.mp[index].Match(newRv, segment)
	}

	if t.cb.OnTraversal != nil {
		return t.cb.OnTraversal(Traversal{
			traveller: t,
			index:     index,
			rv:        rv,
			parentRv:  parentRv,
			key:       key,
			next:      next,
		})
	}

	return next(rv)
}

// Get the length of the path.
func (t Traveller) PathLen() int {
	return len(t.mp)
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
