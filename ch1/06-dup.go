package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// prints: count, text and list of filenames, of lines that appear more than once

func main() {
	counts := make(map[string]int)
	files := make(map[string]map[string]bool)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		// NB, ReadFile returns a byte slice and must therefore converted into a string later. see string(data) in strings.Split()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if files[line] == nil {
				files[line] = make(map[string]bool)
			}
			files[line][filename] = true
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\t", n, line)
			for file, _ := range files[line] {
				fmt.Printf("%s ", file)
			}
			fmt.Println()
		}
	}
}
