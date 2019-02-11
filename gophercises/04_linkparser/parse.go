package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an
// html document
type Link struct {
	Href string
	Text string
}

// Parse parses an HTML document and returns a
// slice of links parsed from it
func Parse(r io.Reader) (links []Link, err error) {
	doc, err := html.Parse(r)
	if err != nil {
		return links, err
	}

	nodes := linkNodes(doc)
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}

	return links, nil
}

func linkNodes(n *html.Node) (nodes []*html.Node) {
	// Check if link node
	if n.Type == html.ElementNode && n.Data == "a" {
		// if it is return it
		return []*html.Node{n}
	}
	// If not search all child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}
	return nodes
}

func buildLink(n *html.Node) (link Link) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	link.Text = text(n)
	return link
}

func text(n *html.Node) string {
	// Base case for text node
	if n.Type == html.TextNode {
		return n.Data
		// Base case for DocumentNode
	} else if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

// Example of DFS
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
