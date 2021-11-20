package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const helloMessage = "Hello, OTUS!"

func main() {
	reversed := stringutil.Reverse(helloMessage)
	fmt.Println(reversed)
}
