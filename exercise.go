package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Modify the echo program to also print os.Args[0], the name of the command that invoked it.

func echo0() {

	fmt.Println(strings.Join(os.Args[0:], " "))
}

// Modify the echo program to print the index and value of each of its arguments, one per line.

func echo1() {

	for idx, arg := range os.Args[1:] {
		fmt.Println(idx)
		fmt.Println(arg)
	}
}

// Experiment to measure the difference in running time between our potentially inefficiet versions and the one that uses strings.Join . [Not completed yet!]

func statusUpdate() string { return "" }

func echo2() {
	// start := time.Now()

	s, sep := "", ""

	fmt.Printf("Time start at %v \n", time.Now().Nanosecond())
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	println(s)

	println(strings.Join(os.Args[1:], " "))

	println("\n")
	fmt.Printf("Time end at %d \n", time.Now().Nanosecond())
}

func main() {
	// echo0()
	// echo1()
	echo2()
}
