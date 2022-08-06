package utils

import "reflect"

func IsEmptyValue(e reflect.Value) bool {
	isEmpty := true
	switch e.Type().Kind() {
	case reflect.String:
		if e.String() != "" {
			isEmpty = false
		}
	case reflect.Array:
		for j := e.Len() - 1; j >= 0; j-- {
			isEmpty = IsEmptyValue(e.Index(j))
			if isEmpty == false {
				break
			}
		}
	case reflect.Float32, reflect.Float64:
		if e.Float() != 0 {
			isEmpty = false
		}
	case reflect.Int32, reflect.Int64:
		if e.Int() != 0 {
			isEmpty = false

		}
	case reflect.Ptr:
		if e.Pointer() != 0 {
			isEmpty = false
		}
	case reflect.Struct:
		for i := e.NumField() - 1; i >= 0; i-- {
			isEmpty = IsEmptyValue(e.Field(i))
			if !isEmpty {
				break
			}
		}
	}

	return isEmpty
}
