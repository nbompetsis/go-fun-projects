package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type MyJSON struct {
	Test  any    `json:"test"`
	Test3 string `json:"test3"`
}

func main() {
	var jsonParsed MyJSON
	err := json.Unmarshal([]byte(`{"test": { "test2": [1, 2, 3] }, "test3" : "..." }`), &jsonParsed)
	if err != nil {
		log.Fatal("error")
	}
	fmt.Printf("%s\n", reflect.TypeOf(jsonParsed))

	switch v := jsonParsed.Test.(type) {
	case map[string]any:
		fmt.Printf("Map found: %v\n", v)
		field1, ok := v["test2"]
		if ok {
			switch v2 := field1.(type) {
			case []any: // slice representation
				fmt.Printf("I found the slice any%v\n", v2)
				for _, v2Element := range v2 {
					fmt.Printf("Type of v2Element %s\n", reflect.TypeOf(v2Element))
					if reflect.TypeOf(v2Element).String() == "float64" {
						fmt.Printf("Float %f\n", v2Element.(float64))
						fmt.Printf("Int %d\n", int(v2Element.(float64)))
					} else {
						fmt.Printf("Didn't recognize v2Element!\n")
					}
				}
			default:
				fmt.Printf("Type not found %s\n", reflect.TypeOf(field1))
			}
		}
	default:
		fmt.Printf("Type not found %s\n", reflect.TypeOf(jsonParsed))
	}
}
