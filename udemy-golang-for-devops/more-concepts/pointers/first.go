package main

import "fmt"

func main() {
	// Pointers
	a := "string"
	testPointer(&a)
	fmt.Printf("a: %s\n", a)

	// Slice
	slice := []string{"1", "2"}
	slice = testSlice(slice)
	fmt.Printf("slice: %v\n", slice)

	testSlice1(&slice)
	fmt.Printf("slice: %v\n", slice)

	// Map
	myMap := make(map[string]string)
	myMap["test"] = "newValue"
	testPointerMap(myMap)
	fmt.Printf("myMap: %v\n", myMap)
}

func testPointer(str *string) {
	*str = "new value"
}

func testSlice(s []string) []string {
	s = append(s, "3")
	return s
}

func testSlice1(s *[]string) {
	*s = append(*s, "4")
}

func testPointerMap(m map[string]string) {
	m["test1"] = "new value"
}
