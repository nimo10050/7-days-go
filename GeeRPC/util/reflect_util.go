package util

import (
	"reflect"
)

func MethodByName(o interface{}, methodName string) reflect.Value {
	v := reflect.ValueOf(o)
	m := v.MethodByName(methodName)
	return m
}
