# htmlsqueeze

htmlsqueeze is a small Go library to extract text out of HTML DOM trees. It is
based on the notions of predicates and extractors. Predicates are rules stating
which nodes are to be extracted when traversing the HTML DOM tree. Extractors
are functions that define how the text is to be extracted from a node.

## TODO

- [ ] implement some more predicates
- [ ] implement some more extractors
- [ ] convenience functions to build up lists of predicate lists

## Examples

Given this HTML page:

```html
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
```

The text content of the `span.important` nodes can be extracted as follows:

```go
doc, _ := html.Parse(strings.NewReader(htmlText))
predicates := [][]htmlsqueeze.Predicate{
    []htmlsqueeze.Predicate{
        htmlsqueeze.TagMatcher("span"),
        htmlsqueeze.ClassMatcher("important"),
    },
}
found := htmlsqueeze.Squeeze(doc, predicates, htmlsqueeze.ExtractChildText)
```

The predicates are given as a list of lists. The top level list contains rules
to be applied to different levels of the tree. The sub-lists contains all the
predicates that a single node must match in order to be extracted.

Given another HTML page:

```html
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
```

The text content of the nodes matchint the CSS selector `div.odd p.yes` can be
extracted as follows:

```go
doc, _ := html.Parse(strings.NewReader(htmlText))
predicates := [][]htmlsqueeze.Predicate{
predicates := [][]htmlsqueeze.Predicate{
    []htmlsqueeze.Predicate{htmlsqueeze.TagMatcher("div"), htmlsqueeze.ClassMatcher("odd")},
    []htmlsqueeze.Predicate{htmlsqueeze.TagMatcher("p"), htmlsqueeze.ClassMatcher("yes")},
}
found := htmlsqueeze.Squeeze(doc, predicates, htmlsqueeze.ExtractChildText)
```
