package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	retstr := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(retstr)
}
