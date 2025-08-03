// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
//
// Exercise 1.4 modifies this program to print the names of all files
// in which each duplicated line occurs.
//
// NOTE: that this exercise adds an error of sorts ... in that if multiple
// files have the same line of text, it doesn't accurately report all cases.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupfiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupfiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, dupfiles[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupfiles map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			dupfiles[input.Text()] = f.Name()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
