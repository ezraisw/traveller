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
	//
	// Match MUST NOT modify the value as it may be overwritten by traversal handlers.
	// In addition, its only purpose is to match for the traversal.
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
		if !wild.Match(m.Pattern, field.Name, m.CaseInsensitive) {
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
		if m.OnlyStringKey {
			if keyRv.Kind() == reflect.String {
				keyStr, ok = keyRv.String(), true
			}
		} else {
			keyStr, ok = AssumeAsString(keyRv)
		}

		if !ok || !wild.Match(m.Pattern, keyStr, m.CaseInsensitive) {
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
		if m.OnlyStringKey {
			continue
		}
		// Force index as string.
		if !wild.Match(m.Pattern, strconv.Itoa(i), m.CaseInsensitive) {
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
	// Whether to always explore using earlier path segments first.
	// This causes traversal order to change but not necessarily the found values.
	StayFirst bool
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

// Needs to be separated so the matcher gets the field value twice.
// Do not use the same reflect.Value.

func (m MatchMulti) op1(childRv reflect.Value, rv reflect.Value, key any, s MatcherSegment) bool {
	if m.StayFirst {
		return s.Stay(childRv, rv, key)
	}
	return s.Next(childRv, rv, key)
}

func (m MatchMulti) op2(childRv reflect.Value, rv reflect.Value, key any, s MatcherSegment) bool {
	if m.StayFirst {
		return s.Next(childRv, rv, key)
	}
	return s.Stay(childRv, rv, key)
}

func (m MatchMulti) matchStruct(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreStruct() {
		return true
	}
	for i := 0; i < rv.NumField(); i++ {
		if field := rv.Type().Field(i); field.IsExported() &&
			(!m.op1(rv.Field(i), rv, field.Name, s) || !m.op2(rv.Field(i), rv, field.Name, s)) {
			return false
		}
	}
	return true
}

func (m MatchMulti) matchMap(rv reflect.Value, s MatcherSegment) bool {
	if s.Traveller().IgnoreMap() {
		return true
	}
	for _, keyRv := range rv.MapKeys() {
		if !m.op1(rv.MapIndex(keyRv), rv, keyRv, s) || !m.op2(rv.MapIndex(keyRv), rv, keyRv, s) {
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
		if !m.op1(rv.Index(i), rv, i, s) || !m.op2(rv.Index(i), rv, i, s) {
			return false
		}
	}
	return true
}
