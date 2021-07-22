package helper

import "reflect"

func Hashable(item interface{}) bool {
	k := reflect.TypeOf(item).Kind()

	return !(k == reflect.Slice ||
		k == reflect.Func ||
		k == reflect.Map)
}
