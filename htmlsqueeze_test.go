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
	found := Squeeze(doc, predicates, ExtractChildText)
	if found[0] != "sit" || found[1] != "consectetur" {
		t.Errorf("expected %v, got %v", []string{"sit", "consectetur"}, found)
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
	predicates := [][]Predicate{
		[]Predicate{TagMatcher("div"), ClassMatcher("odd")},
		[]Predicate{TagMatcher("p"), ClassMatcher("yes")},
	}
	found := Squeeze(doc, predicates, ExtractChildText)
	if got := len(found); got != 4 {
		t.Errorf("expected 4 elements, got %d", got)
	}
	if found[0] != "a" || found[1] != "c" || found[2] != "i" || found[3] != "k" {
		t.Errorf("expected %v, got %v", []string{"a", "c", "i", "k"}, found)
	}
}
