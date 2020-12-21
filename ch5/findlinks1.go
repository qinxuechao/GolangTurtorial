package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


func main() {
	doc, err := html.Parse(os.stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {


}
