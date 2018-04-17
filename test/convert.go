package main

import "fmt"
import "reflect"

func main() {
	var a int = 10
	fmt.Println(reflect.TypeOf(string(a)))
	fmt.Println(float32(a))
}
