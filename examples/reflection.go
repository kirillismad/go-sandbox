package examples

import (
	"fmt"
	"reflect"
)

func DemoReflection() {
	// demoReflectionInt()
	// demoReflectionStruct()
	demoEmptyInterface()
}

func demoReflectionInt() {
	x := 42
	t := reflect.TypeOf(x)

	fmt.Println("t.Name() -", t.Name())
	fmt.Println("t.Kind() -", t.Kind())
	fmt.Println("t.String() -", t.String())
	fmt.Println()

	v := reflect.ValueOf(x)
	fmt.Println("v.Kind() -", v.Kind())
	fmt.Println("v.Int() -", v.Int())
	fmt.Println("v.Type().Name() -", v.Type().Name())
	fmt.Println("v.Type().Kind() -", v.Type().Kind())
}

func demoReflectionStruct() {
	type item struct {
		A float64 `json:"a"`
		B string  `json:"b"`
	}
	x := item{
		A: 43.1,
		B: "Kirill",
	}

	v := reflect.Indirect(reflect.ValueOf(&x))
	fmt.Println("v.Type().Name() -", v.Type().Name())
	fmt.Println("v.Type().Kind() -", v.Type().Kind())
	fmt.Println("v.Type().String() -", v.Type().String())
	fmt.Println("v.Type().NumField() -", v.Type().NumField())
	for i := 0; i < v.Type().NumField(); i++ {
		ft := v.Type().Field(i)
		fmt.Println("\t"+"ft.Name -", ft.Name)
		fmt.Println("\t"+"ft.Tag -", ft.Tag)
		fmt.Println("\t"+"ft.Type.Name() -", ft.Type.Name())
		fmt.Println("\t"+"ft.Type.Kind() -", ft.Type.Kind())

		fv := v.Field(i)
		switch fv.Kind() {
		case reflect.Int:
			fmt.Println("\t"+"fv.Int() -", fv.Int())
		case reflect.String:
			fmt.Println("\t"+"fv.String() -", fv.String())
		default:
			fmt.Println("\t"+"fv.Interface() -", fv.Interface())
		}

		fmt.Println("\t------------------------")
	}
}

func demoEmptyInterface() {
	var o interface{}
	fmt.Println(isEmpty(o))
}

func isEmpty(i interface{}) bool {
	v := reflect.ValueOf(i)

	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func,
		reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
