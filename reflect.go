package goform

import (
	"reflect"
)

func ValueOf(fo FormObject, fieldName string) (r interface{}) {
	v := reflect.ValueOf(fo)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	v = v.FieldByName(fieldName)
	if v.IsValid() {
		return v.Interface()
	}
	v = v.MethodByName(fieldName)
	if v.IsValid() {
		return v.Call([]reflect.Value{})
	}
	return
}
