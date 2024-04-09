package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t1 int = 123
	result := plusOne(t1)
	fmt.Printf("plusOne: %v and (type: %s)\n", result, reflect.TypeOf(result))

	var t2 float64 = 123.12
	result2 := plusOne(t2)
	fmt.Printf("plusOne: %v and (type: %s)\n", result2, reflect.TypeOf(result2))
	result3 := sum(t1, t1)
	fmt.Printf("plusOne: %v and (type: %s)\n", result3, reflect.TypeOf(result3))
}

func plusOne[V int | float64 | int64 | int32 | float32](t V) V {
	return t + 1
}

func sum[V int | float64 | int64 | int32 | float32](v1, v2 V) V {
	return v1 + v2
}
