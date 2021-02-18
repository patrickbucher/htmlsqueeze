package htmlsqueeze

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const dummy = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Test</title>
	</head>
	<body>
		<div class="main">
			<p class="header">Lorem Ipsum</p>
			<p class="content">
				dolor
				<span class="important">sit</span>
				<span class="invisible">amet</span>
				<span class="important">consectetur</span>
			</p>
			<p class="footer">adipiscing</p>
		</div>
	</body>
</html>
`

func TestSqueeze(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(dummy))
	if err != nil {
		t.Error(err)
	}
	predicates := [][]Predicate{
		[]Predicate{
			TagMatcher("span"),
			ClassMatcher("important"),
		},
	}
	found := Squeeze(doc, predicates, ExtractChildText)
	if found[0] != "sit" || found[1] != "consectetur" {
		t.Errorf("expected %v, got %v", []string{"sit", "consectetur"}, found)
	}
}

func TestApply(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(dummy))
	if err != nil {
		t.Error(err)
	}
	predicates := [][]Predicate{[]Predicate{TagMatcher("p")}}
	found := Apply(doc, predicates)
	if got := len(found); got != 3 {
		t.Errorf("expected 3 nodes, got %d", got)
	}
}
