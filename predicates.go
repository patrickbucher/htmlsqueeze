package htmlsqueeze

import (
	"strings"

	"golang.org/x/net/html"
)

// Predicate returns true if the given node satisfies a condition, and false
// otherwise.
type Predicate func(n *html.Node) bool

// TagMatcher creates a predicate that tests if a node is an element node of
// the given name.
func TagMatcher(name string) Predicate {
	return func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		return n.Data == name
	}
}

// ClassMatcher creates a predicate that tests if a node has a class attribute
// containing the given name.
func ClassMatcher(name string) Predicate {
	return func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, attr := range n.Attr {
			if attr.Key != "class" {
				continue
			}
			for _, class := range strings.Fields(attr.Val) {
				if class == name {
					return true
				}
			}
		}
		return false
	}
}
