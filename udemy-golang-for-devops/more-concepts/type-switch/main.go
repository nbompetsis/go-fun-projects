package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t string = "this is a string"
	var t2 *string = &t
	discoverType(t)
	discoverType(t2)
	var t3 int = 123
	discoverType(t3)
	discoverType(nil)
}

func discoverType(t any) {
	switch v := t.(type) {
	case string:
		t2 := v + "..."
		fmt.Printf("String found: %s\n", t2)
	case *string:
		fmt.Printf("Ptr string found: %s\n", *v)
	case int:
		fmt.Printf("Integer found: %d\n", v)
	default:
		objType := reflect.TypeOf(t)
		if objType != nil {
			fmt.Printf("Type not found %s\n", objType)
		} else {
			fmt.Printf("Type is nil\n")
		}

	}

}
