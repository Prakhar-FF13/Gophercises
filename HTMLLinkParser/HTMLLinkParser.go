package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	dfs(doc, "")
	fmt.Printf("%+v\n", doc.FirstChild)
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}

var exampleHTML = `
<a href="/dog">
	<span>
		Something in a span
	</span>
	Text not in a span
	<b>
		Bold text!
	</b>
</a>
`

func main() {
	r := strings.NewReader(exampleHTML)
	Links, err := Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", Links)
}
