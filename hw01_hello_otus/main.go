package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	var tempString string = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(tempString))
}
