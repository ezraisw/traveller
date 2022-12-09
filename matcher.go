package traveller

import (
	"reflect"
	"strconv"

	"github.com/gertd/wild"
)

type Matcher interface {
	// Match the given reflect.Value. Return true to continue traversing.
	//
	// Call MatcherSegment.Next to continue traversal using the next path segment or
	// MatcherSegment.Stay to keep using the current path segment for a new value.
	Match(reflect.Value, MatcherSegment) (keepSearching bool)
}

// Exact match by value.
type MatchExact struct {
	// The value to match with.
	//
	// To match a field name of a struct, use a string.
	// To match an index of an array/slice, use an int.
	// To match a map key, use the correct key type of that map.
	Value any
}

// Compile-time implementation check.
var _ Matcher = (*MatchExact)(nil)

func (m MatchExact) Match(rv reflect.Value, s MatcherSegment) bool {
	switch rv := Unbox(rv); rv.Kind() {
	case reflect.Struct:
		return m.matchStruct(rv, s)
	case reflect.Map:
		return m.matchMap(rv, s)
	case reflect.Array, reflect.Slice:
		return m.matchArray(rv, s)
	}
	return true
}

func (m MatchExact) matchStruct(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreStruct() {
		return true
	}
	if name, ok := m.Value.(string); ok {
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := rt.Field(i)
			if !field.IsExported() {
				continue
			}
			fieldRv := rv.Field(i)
			if field.Name == name && !s.Next(fieldRv, rv, name) {
				return false
			}
			// Check embedded values.
			if !s.Traveller().NoFlatEmbeds() && field.Anonymous && !s.Stay(fieldRv, rv, name) {
				return false
			}
		}
	}
	return true
}

func (m MatchExact) matchMap(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreMap() {
		return true
	}
	keyRv := reflect.ValueOf(m.Value)
	if valueRv := rv.MapIndex(keyRv); valueRv.IsValid() {
		if !s.Next(valueRv, rv, keyRv) {
			return false
		}
	}
	return true
}

func (m MatchExact) matchArray(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreArray() {
		return true
	}
	if i, ok := m.Value.(int); ok && i >= 0 && i < rv.Len() {
		if !s.Next(rv.Index(i), rv, i) {
			return false
		}
	}
	return true
}

// Match by a wildcard string pattern.
type MatchPattern struct {
	// The string pattern to use. Uses wildcard pattern.
	Pattern string

	// Extra options.
	Options MatchPatternOptions
}

type MatchPatternOptions struct {
	// Only try to match keys that are a type of string.
	// If false, attempt to convert non string keys into a string.
	OnlyStringKey bool

	// Disregard letter cases.
	// If true, strings like "JoHn" and "john" will be the same.
	CaseInsensitive bool
}

// Compile-time implementation check.
var _ Matcher = (*MatchPattern)(nil)

func (m MatchPattern) Match(rv reflect.Value, s MatcherSegment) bool {
	switch rv := Unbox(rv); rv.Kind() {
	case reflect.Struct:
		return m.matchStruct(rv, s)
	case reflect.Map:
		return m.matchMap(rv, s)
	case reflect.Array, reflect.Slice:
		return m.matchArray(rv, s)
	}
	return true
}

func (m MatchPattern) matchStruct(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreStruct() {
		return true
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		if !field.IsExported() {
			continue
		}
		if !wild.Match(m.Pattern, field.Name, m.Options.CaseInsensitive) {
			continue
		}
		fieldRv := rv.Field(i)
		if !s.Next(fieldRv, rv, field.Name) {
			return false
		}
		// Check embedded values.
		if !s.Traveller().NoFlatEmbeds() && field.Anonymous && !s.Stay(fieldRv, rv, field.Name) {
			return false
		}
	}
	return true
}

func (m MatchPattern) matchMap(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreMap() {
		return true
	}
	for it := rv.MapRange(); it.Next(); {
		keyRv := it.Key()
		// Force key as string.
		var (
			keyStr string
			ok     bool
		)
		if m.Options.OnlyStringKey {
			if keyRv.Kind() == reflect.String {
				keyStr, ok = keyRv.String(), true
			}
		} else {
			keyStr, ok = AssumeAsString(keyRv)
		}

		if !ok || !wild.Match(m.Pattern, keyStr, m.Options.CaseInsensitive) {
			continue
		}
		if !s.Next(it.Value(), rv, keyRv) {
			return false
		}
	}
	return true
}

func (m MatchPattern) matchArray(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreArray() {
		return true
	}
	for i := 0; i < rv.Len(); i++ {
		// Array indexes are ints, therefore it is inevitable when OnlyStringKey is active.
		if m.Options.OnlyStringKey {
			continue
		}
		// Force index as string.
		if !wild.Match(m.Pattern, strconv.Itoa(i), m.Options.CaseInsensitive) {
			continue
		}
		if !s.Next(rv.Index(i), rv, i) {
			return false
		}
	}
	return true
}

// Recursive free matcher.
type MatchMulti struct {
}

// Compile-time implementation check.
var _ Matcher = (*MatchMulti)(nil)

func (m MatchMulti) Match(rv reflect.Value, s MatcherSegment) bool {
	switch rv := Unbox(rv); rv.Kind() {
	case reflect.Struct:
		return m.matchStruct(rv, s)
	case reflect.Map:
		return m.matchMap(rv, s)
	case reflect.Array, reflect.Slice:
		return m.matchArray(rv, s)
	}
	return true
}

func (m MatchMulti) matchStruct(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreStruct() {
		return true
	}
	for i := 0; i < rv.NumField(); i++ {
		if field, fieldRv := rv.Type().Field(i), rv.Field(i); field.IsExported() &&
			(!s.Stay(fieldRv, rv, field.Name) || !s.Next(fieldRv, rv, field.Name)) {
			return false
		}
	}
	return true
}

func (m MatchMulti) matchMap(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreMap() {
		return true
	}
	for it := rv.MapRange(); it.Next(); {
		if keyRv, valueRv := it.Key(), it.Value(); !s.Stay(valueRv, rv, keyRv) || !s.Next(valueRv, rv, keyRv) {
			return false
		}
	}
	return true
}

func (m MatchMulti) matchArray(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreArray() {
		return true
	}
	for i := 0; i < rv.Len(); i++ {
		if elRv := rv.Index(i); !s.Stay(elRv, rv, i) || !s.Next(elRv, rv, i) {
			return false
		}
	}
	return true
}
