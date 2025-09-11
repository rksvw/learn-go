// Print the name of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	myFiles := make(map[string]map[string]bool)

	if len(files) == 0 {
		countLines(os.Stdin, counts, myFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, myFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			file := []string{}
			for fname := range myFiles[line] {
				file = append(file, fname)
			}

			fmt.Printf("Line: %q, Count: %d, Files: %v\n", line, n, file)
		}
	}
}

func countLines(f *os.File, counts map[string]int, myFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		counts[line]++

		if myFiles[line] == nil {
			myFiles[line] = make(map[string]bool)
		}

		myFiles[line][f.Name()] = true
	}
}
