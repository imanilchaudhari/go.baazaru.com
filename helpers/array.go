package helpers

import (
	"fmt"
	"reflect"
)

// Keys creates an array of the own enumerable map keys or struct field names.
func Keys(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()

	if value.Kind() == reflect.Map {
		keys := value.MapKeys()

		length := len(keys)

		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Key()), length, length)

		for i, key := range keys {
			resultSlice.Index(i).Set(key)
		}

		return resultSlice.Interface()
	}

	if value.Kind() == reflect.Struct {
		length := value.NumField()

		resultSlice := make([]string, length)

		for i := 0; i < length; i++ {
			resultSlice[i] = valueType.Field(i).Name
		}

		return resultSlice
	}

	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

// Values creates an array of the own enumerable map values or struct field values.
func Values(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()

	if value.Kind() == reflect.Map {
		keys := value.MapKeys()

		length := len(keys)

		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), length, length)

		for i, key := range keys {
			resultSlice.Index(i).Set(value.MapIndex(key))
		}

		return resultSlice.Interface()
	}

	if value.Kind() == reflect.Struct {
		length := value.NumField()

		resultSlice := make([]interface{}, length)

		for i := 0; i < length; i++ {
			resultSlice[i] = value.Field(i).Interface()
		}

		return resultSlice
	}

	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

// StringInSlice verifies if a string is present in a slice
func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// DiffSlices takes two integer slices and returns the diff between them
// e.g. DiffSlices(X, Y) will return the elements that are only in X
// If an integer is present in both slices but not the same number of time,
// it will be reflected in the result
func DiffSlices(X, Y []int) (ret []int) {
	m := make(map[int]int)

	for _, y := range Y {
		m[y]++
	}

	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}
