package htmlsqueeze

import (
	"golang.org/x/net/html"
)

// Squeeze applies the given predicates to the given node, applies the given
// extractor to the matching nodes, and returns the extracted text as a list of
// strings.
func Squeeze(n *html.Node, predicates [][]Predicate, extract Extractor) []string {
	texts := make([]string, 0)
	nodes := Apply(n, predicates)
	for _, node := range nodes {
		texts = append(texts, extract(node))
	}
	return texts
}

// Apply applies the given predicates to the given node, and returns the
// matching nodes. If a node satisfies all the predicates of the first
// sub-list, the remaining predicates are applied to the node's children;
// otherwise all the predicates are applied to the node's children. If no more
// predicates are left to be satisfied, the node is considered a match and
// returned.
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

// MatchAll applies the given predicates to a node, returns true if the node
// satisfies all those predicates, and false otherwise.
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
