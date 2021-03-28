package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var x float64 = 3.4
	// fmt.Println("type:", reflect.TypeOf(x))
	// fmt.Println("value:", reflect.ValueOf(x).String())

	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	// var x uint8 = 'x'
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8)
	// x = uint8(v.Uint())
	// fmt.Println(x)

	// type MyInt int
	// var x MyInt = 7
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is:", v.Kind())
	// fmt.Println("value:", v.String())

	// v := reflect.ValueOf(3.14)
	// y := v.Interface().(float64)
	// fmt.Println(y)
	// fmt.Println(v.Interface())
	// fmt.Printf("value is %7.1e\n", v.Interface())

	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("settability of v:", v.CanSet())

	// var x float64 = 3.4
	// p := reflect.ValueOf(&x)
	// fmt.Println("type of p:", p.Type())
	// fmt.Println("settability of p:", p.CanSet())
	// v := p.Elem()
	// fmt.Println("settability of v:", v.CanSet())
	// v.SetFloat(7.1)
	// fmt.Println(v.Interface())
	// fmt.Println(x)

	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
