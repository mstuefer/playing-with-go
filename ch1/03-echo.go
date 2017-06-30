// print index and argument of command, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i, arg)
	}
}
