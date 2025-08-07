package main

import "fmt"

func assert(condition bool, msg string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, args...))
	}
}

func if_else[T any](condition bool, true_val T, false_val T) T {
	if condition {
		return true_val
	} else {
		return false_val
	}
}
