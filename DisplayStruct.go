package main

import (
	"fmt"
	"reflect"
)

// DisplayStruct prints all exported fields of any struct in a clean format.
func DisplayStruct(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// Handle pointer to struct
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("DisplayStruct: not a struct")
		return
	}

	fmt.Printf("Struct: %s\n", typ.Name())
	fmt.Println("-------------------------")
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Only show exported fields
		if field.PkgPath == "" {
			fmt.Printf("%-20s : %v\n", field.Name, value.Interface())
		}
	}
	fmt.Println()
}
