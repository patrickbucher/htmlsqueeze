package htmlsqueeze

import "golang.org/x/net/html"

func ExtractChildText(n *html.Node) string {
	if n.FirstChild.Type == html.TextNode {
		return n.FirstChild.Data
	}
	return ""
}
