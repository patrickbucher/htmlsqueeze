package htmlsqueeze

import (
	"golang.org/x/net/html"
)

type Predicate func(n *html.Node) bool

type Extractor func(n *html.Node) string

func Squeeze(n *html.Node, predicates [][]Predicate, extract Extractor) []string {
	texts := make([]string, 0)
	nodes := Apply(n, predicates)
	for _, node := range nodes {
		texts = append(texts, extract(node))
	}
	return texts
}

func Apply(n *html.Node, predicates [][]Predicate) []*html.Node {
	if len(predicates) == 0 {
		return []*html.Node{n}
	}
	nodes := make([]*html.Node, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if MatchAll(c, predicates[0]) {
			nodes = append(nodes, Apply(c, predicates[1:])...)
		} else {
			nodes = append(nodes, Apply(c, predicates)...)
		}
	}
	return nodes
}

func MatchAll(n *html.Node, predicates []Predicate) bool {
	if len(predicates) == 0 {
		return true
	}
	for _, predicate := range predicates {
		if !predicate(n) {
			return false
		}
	}
	return true
}
