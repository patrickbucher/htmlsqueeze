package htmlsqueeze

import (
	"strings"

	"golang.org/x/net/html"
)

// Predicate returns true if the given node satisfies a condition, and false
// otherwise.
type Predicate func(n *html.Node) bool

// DontMatch is a dummy matcher that never matches.
func DontMatch() Predicate {
	return func(n *html.Node) bool {
		return false
	}
}

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

// TagClassMatchersOf expects multiple selectors encoded in a single string,
// such as "div.main p.text", which are split into fields, and given as a
// parameter to TagClassMatchers to produce the specified predicates.
func TagClassMatchersOf(selectors string) [][]Predicate {
	return TagClassMatchers(strings.Fields(selectors))
}

// TagClassMatchers expects multiple selectors. For each selector,
// TagClassMatcher is invoked to produce a list of predicates. Those predicates
// are returned as a list of predicate lists, which can be used for the Squeeze
// and Apply functions.
func TagClassMatchers(selectors []string) [][]Predicate {
	predicates := make([][]Predicate, 0)
	for _, selector := range selectors {
		predicates = append(predicates, TagClassMatcher(selector))
	}
	return predicates
}

// TagClassMatcher expects a selector like "div.main" (matching div elements
// with class main) or "div" (just matching div elements) and produces a list
// of according predicates.
// If the selector is malformed, a predicate list containing DontMatch is
// returned, which matches to nothing.
func TagClassMatcher(selector string) []Predicate {
	if strings.Contains(selector, ".") {
		// match element and class
		parts := strings.Split(selector, ".")
		if len(parts) != 2 {
			return []Predicate{DontMatch()}
		}
		return []Predicate{
			TagMatcher(parts[0]),
			ClassMatcher(parts[1]),
		}
	}
	// just match on the element
	return []Predicate{TagMatcher(selector)}
}
