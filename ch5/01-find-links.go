// get all links from a html-file passed via stdin
// including the ones in images and style sheets

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "00-find-links: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			links = collectLinks(n.Attr, "href", links)
		case "img":
			links = collectLinks(n.Attr, "src", links)
		case "link":
			links = collectLinks(n.Attr, "href", links)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func collectLinks(attributes []html.Attribute, attributeName string, links []string) []string {
	for _, a := range attributes {
		if a.Key == attributeName {
			links = append(links, a.Val)
		}
	}
	return links
}
