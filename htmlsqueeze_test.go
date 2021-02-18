package htmlsqueeze

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const loremIpsumHTML = `
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
				<span class="important">sit amet</span>
				<span class="invisible">consectetur</span>
				<span class="important">adipiscing <br /> elit</span>
			</p>
			<p class="footer">nunc</p>
		</div>
	</body>
</html>
`

func TestSqueeze(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(loremIpsumHTML))
	if err != nil {
		t.Error(err)
	}
	predicates := [][]Predicate{
		[]Predicate{
			TagMatcher("span"),
			ClassMatcher("important"),
		},
	}
	found := Squeeze(doc, predicates, ExtractChildrenTexts)
	if found[0] != "sit amet" || found[1] != "adipiscing elit" {
		t.Errorf("expected %v, got %v", []string{"sit amet", "adipiscing elit"}, found)
	}
}

func TestApply(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(loremIpsumHTML))
	if err != nil {
		t.Error(err)
	}
	predicates := [][]Predicate{[]Predicate{TagMatcher("p")}}
	found := Apply(doc, predicates)
	if got := len(found); got != 3 {
		t.Errorf("expected 3 nodes, got %d", got)
	}
}

const subTreeHTML = `
<div class="main">
	<div class="odd">
		<p class="yes">a</p>
		<p class="no">b</p>
		<p class="yes">c</p>
		<p class="no">d</p>
	</div>
	<div class="even">
		<p class="yes">e</p>
		<p class="no">f</p>
		<p class="yes">g</p>
		<p class="no">h</p>
	</div>
	<div class="odd">
		<p class="yes">i</p>
		<p class="no">j</p>
		<p class="yes">k</p>
		<p class="no">l</p>
	</div>
	<div class="even">
		<p class="yes">m</p>
		<p class="no">n</p>
		<p class="yes">o</p>
		<p class="no">p</p>
	</div>
</div>
`

func TestSqueezeSubTrees(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(subTreeHTML))
	if err != nil {
		t.Error(err)
	}
	found := SqueezeSelector(doc, "div.odd p.yes", ExtractChildText)
	if got := len(found); got != 4 {
		t.Errorf("expected 4 elements, got %d", got)
	}
	if found[0] != "a" || found[1] != "c" || found[2] != "i" || found[3] != "k" {
		t.Errorf("expected %v, got %v", []string{"a", "c", "i", "k"}, found)
	}
}

func TestTagClassMatchersOf(t *testing.T) {
	predicates := TagClassMatchersOf("div.main div p.text span.important")
	if len(predicates) != 4 {
		t.Errorf("expected 4 predicate lists, got %d", len(predicates))
	}
	if len(predicates[0]) != 2 || len(predicates[1]) != 1 ||
		len(predicates[2]) != 2 || len(predicates[3]) != 2 {
		t.Error("predicate sub-lists have the wrong length")
	}
}
