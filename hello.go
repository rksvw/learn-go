package main

import (
	"fmt"
	"os"
	"strings"
)

func identifier() {
	// Short variable declaration, used only within a function (not for package-level variables).
	s := "Hello"

	// Default initialization to the zero value of strings -> ""
	var st string

	// Rarely used except when declaring multiple variables
	var str = "Hello str"

	// explicit about the variable's tpe
	var istr string = " Hello iStr"

	st = "Hello st of string"

	println(s, st, str, istr)

	fmt.Println(strings.Join(os.Args[1:], " "))
}

func arg() {
	fmt.Println(os.Args[1:])
}

func main() {
	// identifier()
	arg()
}
