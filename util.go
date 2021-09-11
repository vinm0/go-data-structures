package main

import "reflect"

func Hashable(item interface{}) bool {
	if item == nil {
		return false
	}

	k := reflect.TypeOf(item).Kind()

	return !(k == reflect.Slice ||
		k == reflect.Func ||
		k == reflect.Map)
}

func ValidIndex(index int, size int) bool {
	return index >= 0 && index < size
}
