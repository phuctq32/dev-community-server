package common

import (
	"fmt"
	"reflect"
)

// StructToMap: convert struct to Map[string]interface{}
func StructToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // Non-structural return error
		return nil, fmt.Errorf("StructToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if tagValue := field.Tag.Get("map"); tagValue != "-" {
			value := reflect.ValueOf(v.Field(i).Interface())
			if value.Kind() == reflect.Ptr {
				if value.IsNil() {
					continue
				}
				value = value.Elem()
			} else if value.IsZero() {
				continue
			}

			// if value is struct, using recursion to convert
			if value.Kind() == reflect.Struct {
				convertedVal, err := StructToMap(value.Interface())
				if err != nil {
					return nil, err
				}
				out[tagValue] = convertedVal
			} else {
				out[tagValue] = value.Interface()
			}
		}
	}
	return out, nil
}
