package main

import (
	"fmt"
	"strings"
)

func main() {
	var broken string = "Brok#n T#xt"
	var replacer *strings.Replacer = strings.NewReplacer("#", "e")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
