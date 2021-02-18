package htmlsqueeze

import (
	"strings"

	"golang.org/x/net/html"
)

func TagMatcher(name string) Predicate {
	return func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		return n.Data == name
	}
}

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
