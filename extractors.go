package htmlsqueeze

import (
	"strings"

	"golang.org/x/net/html"
)

// Extractor is a function that extracts text of a node.
type Extractor func(n *html.Node) string

// ExtractChildText returns the text of n's first child, if it is a text node,
// and the empty string otherwise.
func ExtractChildText(n *html.Node) string {
	if n.FirstChild.Type == html.TextNode {
		return n.FirstChild.Data
	}
	return ""
}

// ExtractChildrenTexts returns the text of n's children that are text nodes
// separated by space.
func ExtractChildrenTexts(n *html.Node) string {
	texts := make([]string, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.TextNode {
			continue
		}
		texts = append(texts, strings.TrimSpace(c.Data))
	}
	return strings.Join(texts, " ")
}
