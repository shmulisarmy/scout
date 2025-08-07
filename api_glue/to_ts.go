package apiglue

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Maps Go reflect.Type to corresponding TypeScript type as string
func typeToTSType(t reflect.Type, queue *[]reflect.Type, parsed map[string]bool) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.Slice, reflect.Array:
		return typeToTSType(t.Elem(), queue, parsed) + "[]"
	case reflect.Map:
		keyType := typeToTSType(t.Key(), queue, parsed)
		valueType := typeToTSType(t.Elem(), queue, parsed)
		return "{ [key: " + keyType + "]: " + valueType + " }"
	case reflect.Struct:
		if t.Name() != "" && !parsed[t.Name()] {
			*queue = append(*queue, t)
		}
		return t.Name()
	default:
		return "any"
	}
}

// var interface_mappings = map[string][]reflect.Type{
// 	"Spot": []reflect.Type{
// 		reflect.TypeOf(Property{}),
// 		reflect.TypeOf(Get_taxed{}),
// 		reflect.TypeOf(Recieve_payment{}),
// 	},
// }

// Converts a Go struct type into a TypeScript type definition string
func fromStructToTSType(t reflect.Type, queue *[]reflect.Type, parsed map[string]bool) string {
	if t.Kind() == reflect.Array {
		return typeToTSType(t.Elem(), queue, parsed) + "[]"
	}

	//handles struct case
	typeName := t.Name()
	// if interface_mappings[typeName] != nil {
	// 	union_name := Map(interface_mappings[typeName], func(t reflect.Type) string {
	// 		return typeToTSType(t, queue, parsed)
	// 	})
	// 	return strings.Join(union_name, "|")
	// }
	if parsed[typeName] || typeName == "" {
		return ""
	}
	parsed[typeName] = true

	res := "export type " + typeName + " = {\n"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonName := field.Tag.Get("json")
		if jsonName == "" {
			jsonName = strings.ToLower(field.Name)
		} else {
			jsonName = strings.Split(jsonName, ",")[0]
		}
		tsType := typeToTSType(field.Type, queue, parsed)
		res += "  " + jsonName + ": " + tsType + ";\n"
	}
	res += "}\n"
	return res
}

// === Example structs ===

type Ts_Type_Converter struct {
	parsed map[string]bool
	queue  []reflect.Type
	file   string
}

func (this *Ts_Type_Converter) Convert() string {

	var results strings.Builder

	for len(this.queue) > 0 {
		t := this.queue[0]
		this.queue = this.queue[1:]

		ts := fromStructToTSType(t, &this.queue, this.parsed)
		if ts != "" {
			results.WriteString(ts + "\n")
		}
	}

	// Write the resultsts to a file
	err := os.WriteFile(this.file, []byte(results.String()), 0644)
	if err != nil {
		fmt.Println("Failed to write "+this.file+":", err)
	} else {
		fmt.Println("Generated " + this.file + " successfully.")
	}
	return results.String()
}
