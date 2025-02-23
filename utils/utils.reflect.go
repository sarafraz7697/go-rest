package utils

import (
	"reflect"
	"strings"
)

// MapStruct maps fields from src (DTO) to dst (Model) using reflection.
func MapStruct(src, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem() // Ensure dst is a pointer

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i).Name
		srcField := srcVal.Field(i)

		dstField := dstVal.FieldByName(field)

		// Only set if field exists and is assignable
		if dstField.IsValid() && dstField.CanSet() && srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		}
	}
}

// StructToMap converts a struct to map[string]interface{}, using JSON tags as keys.
func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(input)

	// Ensure we're working with a struct pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Get JSON tag, default to field name if empty
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Handle struct fields (only take first tag part before ',')
		jsonKey := strings.Split(jsonTag, ",")[0]
		result[jsonKey] = fieldValue.Interface()
	}

	return result
}
