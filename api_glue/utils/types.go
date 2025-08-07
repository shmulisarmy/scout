package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func FirstCharOppositeCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.Title(s)
}

func CreateNameFromType[T any](t T) string {
	typ := reflect.TypeOf(t)
	fmt.Println("reflect.TypeOf(t).Kind()", typ.Kind())

	switch typ.Kind() {
	case reflect.Slice:
		elem := typ.Elem()
		return FirstCharOppositeCase(elem.Name()) + "s"

	case reflect.Map:
		key := typ.Key()
		value := typ.Elem()
		keyName := FirstCharOppositeCase(key.Name())
		valueName := FirstCharOppositeCase(value.Name())
		return keyName + "_to_" + valueName + "_map"

	default:
		return FirstCharOppositeCase(typ.Name())
	}
}
