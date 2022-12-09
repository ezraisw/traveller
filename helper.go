package traveller

import (
	"fmt"
	"reflect"
	"strconv"
)

// This is similar to reflect.Indirect, but unboxes the interface type
// first before exposing the actual type behind the pointer.
func Unbox(rv reflect.Value) reflect.Value {
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return reflect.Indirect(rv)
}

// Assumes the given reflect value as a string.
//
// The type int, uint, float, complex, and bool gets formatted into string.
// The value that implements fmt.Stringer or error will use the corresponding method.
func AssumeAsString(rv reflect.Value) (str string, ok bool) {
	switch rv.Kind() {
	case reflect.String:
		str, ok = rv.String(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str, ok = strconv.FormatInt(rv.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str, ok = strconv.FormatUint(rv.Uint(), 10), true
	case reflect.Float32, reflect.Float64:
		str, ok = strconv.FormatFloat(rv.Float(), 'e', -1, rv.Type().Bits()), true
	case reflect.Complex64, reflect.Complex128:
		str, ok = strconv.FormatComplex(rv.Complex(), 'e', -1, rv.Type().Bits()), true
	case reflect.Bool:
		str, ok = strconv.FormatBool(rv.Bool()), true
	default:
		switch v := rv.Interface().(type) {
		case error:
			str, ok = v.Error(), true
		case fmt.Stringer:
			str, ok = v.String(), true
		}
	}
	return
}
