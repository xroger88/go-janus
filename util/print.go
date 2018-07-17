package util

import (
	"fmt"
	"reflect"
)

func PrintValue(depth int, any interface{}) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "\t"
	}
	va := reflect.ValueOf(any)
	if va.Kind() == reflect.Interface || va.Kind() == reflect.Ptr {
		va = va.Elem()
	}
	ty := va.Type()

	if ty.Kind() != reflect.Struct {
		fmt.Printf("%s%s = %v\n", indent, ty.Name(), va.Interface())
		return
	}

	for i := 0; i < va.NumField(); i++ {
		ft := ty.Field(i) // get struct i-th field type information
		fv := va.Field(i)
		if fv.Type().Kind() == reflect.Struct {
			fmt.Printf("%s%s = \n", indent, ft.Name)
			PrintValue(depth+1, fv.Interface())
		} else {
			if fv.Type().Kind() == reflect.String {
				fmt.Printf("%s%s = %q\n", indent, ft.Name, fv.Interface())
			} else {
				fmt.Printf("%s%s = %v\n", indent, ft.Name, fv.Interface())
			}
		}
	}
}
