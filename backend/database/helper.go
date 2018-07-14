package database

import "reflect"

func getElemType(i interface{}) reflect.Type {
	for t := reflect.TypeOf(i); ; {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice:
			t = t.Elem()
		default:
			return t
		}
	}
}
