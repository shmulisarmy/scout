package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// DisplayStructDetailed returns a string representation of any struct,
// showing all exported fields, their types, values, and nested structs.
func DisplayStructDetailed(v interface{}) string {
	var b strings.Builder
	writeStructRecursive(&b, reflect.ValueOf(v), 0)
	return b.String()
}

func writeStructRecursive(b *strings.Builder, val reflect.Value, indent int) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		b.WriteString(fmt.Sprintf("%s(not a struct)\n", strings.Repeat("  ", indent)))
		return
	}

	typ := val.Type()
	indentStr := strings.Repeat("  ", indent)
	b.WriteString(fmt.Sprintf("%sStruct: %s\n", indentStr, typ.Name()))

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		fieldType := field.Type
		fieldName := field.Name

		if fieldType.Kind() == reflect.Struct && fieldType != reflect.TypeOf(time.Time{}) {
			b.WriteString(fmt.Sprintf("%s  %s (%s):\n", indentStr, fieldName, fieldType))
			writeStructRecursive(b, fieldVal, indent+2)
		} else {
			b.WriteString(fmt.Sprintf("%s  %s (%s): %v\n", indentStr, fieldName, fieldType, fieldVal.Interface()))
		}
	}
}
