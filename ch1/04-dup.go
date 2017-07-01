// dup prints each line that appears more than once in the std input

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// a map with string as key and int as value. NB the key of a map, has to be of a type whos value can be compared via ==
	// build-in function 'make' just creates the new map
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	// each call to input.Scan() reads a line and removes newline char at and
	// input.Text() can be used to retrieve that scanned value
	// as long as there is a line, input.Scan() returns true, otherwise false

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
