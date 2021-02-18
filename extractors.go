package htmlsqueeze

import "golang.org/x/net/html"

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
