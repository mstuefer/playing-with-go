// fetches webcontent of given urls in parallel, reporting size and time

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // creates a channel of strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}

	// receive and print all lines sent to the channel ch
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // sends error to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // sends a summary line on the channel ch
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%7d bytes in %.2f seconds for %s", nbytes, secs, url)
}
