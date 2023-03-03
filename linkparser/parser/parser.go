package parser

import (
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var ret []Link
	for _, n := range linkNodes(doc) {
		ret = append(ret, buildLink(n))
	}
	return ret, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var res []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, linkNodes(c)...)
	}
	return res
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
			break
		}
	}
	ret.Text = "Todo href text"

	return ret
}
