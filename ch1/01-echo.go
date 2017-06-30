package main

import (
	"fmt"
	"os"
) // import n packages as a parenthesized list

func main() {
	var s, sep string                   // declares two variables, s and sep of type string
	for i := 1; i < len(os.Args); i++ { // := is a 'short variable declaration' means it gives them an appropriate type based on the initialiser value
		s += sep + os.Args[i]
		// NB: when sep was declared, it does get initialised with its default 'zero value' which on a string is the empty string. Therefore we get the first sep as empty string, and afterwards we assign it a space, so all passed arguments are divided by a single space
		sep = " "
	}
	fmt.Println(s)
}
