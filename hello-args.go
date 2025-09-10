// Echo1 prints its comman-line argumnets.
package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	// fmt.Println(s)

	echo2()
}

func echo2() {
	// := short variable declaration ( used to automatic detect the data type based on value)
	s, sep := "", ""

	// Range := Produce a pair of values: index and value of the element at that index
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	// _ : blank identifier ( require variable name but not logic )

	fmt.Println(s)
}
