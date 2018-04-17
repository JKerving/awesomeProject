package main

import "reflect"

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	println("type:", v.Type())
	println("kind is float64:", v.Kind() == reflect.Float64)
	println("value:",v.Float())
}
