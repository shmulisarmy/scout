package jsonutils

import (
	"encoding/json"
	"reflect"
)

func Compare(a any, b any) bool {
	// Handle nil cases
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Convert both values to JSON bytes
	jsonA, errA := json.Marshal(a)
	if errA != nil {
		return false
	}

	jsonB, errB := json.Marshal(b)
	if errB != nil {
		return false
	}

	// Parse JSON back to interface{} to normalize the structure
	var normalizedA, normalizedB interface{}

	if err := json.Unmarshal(jsonA, &normalizedA); err != nil {
		return false
	}

	if err := json.Unmarshal(jsonB, &normalizedB); err != nil {
		return false
	}

	// Use reflect.DeepEqual on the normalized structures
	return reflect.DeepEqual(normalizedA, normalizedB)
}

func Copy[T any](src *T) (*T, error) {
	// Handle nil case
	if src == nil {
		return nil, nil
	}

	// Marshal to JSON bytes
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	// Create a new instance of T
	var dst T
	err = json.Unmarshal(jsonBytes, &dst)
	if err != nil {
		return nil, err
	}

	return &dst, nil
}

// Alternative version that panics on error (similar to must-style functions)
func Json_copy_must[T any](src *T) *T {
	dst, err := Copy(src)
	if err != nil {
		panic(err)
	}
	return dst
}
