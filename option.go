package traveller

// Represents an optional setting for Traveller.
type TravellerOption func(*Traveller)

// Prevents flattening of embedded types into the matcher.
// Embedded structs will be only be treated as normal fields.
func WithNoFlatEmbeds(noFlatEmbeds bool) TravellerOption {
	return func(t *Traveller) {
		t.noFlatEmbeds = noFlatEmbeds
	}
}

// Prevent traversal into structs.
// Does not prevent returns of struct values.
func WithIgnoreStruct(ignoreStruct bool) TravellerOption {
	return func(t *Traveller) {
		t.ignoreStruct = ignoreStruct
	}
}

// Prevent traversal into maps.
// Does not prevent returns of map values.
func WithIgnoreMap(ignoreMap bool) TravellerOption {
	return func(t *Traveller) {
		t.ignoreMap = ignoreMap
	}
}

// Prevent traversal into arrays.
// Does not prevent returns of array values.
func WithIgnoreArray(ignoreArray bool) TravellerOption {
	return func(t *Traveller) {
		t.ignoreArray = ignoreArray
	}
}
