// inspired by curl, fetchUrl prints the uninterpreted content found at a URL

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // makes an HTTP request and, if there is no error, returns the result in the response struct 'resp'
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchUrl: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n\nStatus: %v\n\n", resp.Status)
		// the Body field of the resp struct contains the server response as a readable stream
		_, cerr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if cerr != nil {
			fmt.Fprintf(os.Stderr, "fetchUrl: reading %s: %v\n", url, cerr)
			os.Exit(1)
		}
	}
}
