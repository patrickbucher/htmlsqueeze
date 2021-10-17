# htmlsqueeze

htmlsqueeze is a small Go library to extract text out of HTML DOM trees. It is
based on the notions of predicates and extractors. Predicates are rules stating
which nodes are to be extracted when traversing the HTML DOM tree. Extractors
are functions that define how the text is to be extracted from a node.

## TODO

- [ ] implement some more predicates
- [ ] implement some more extractors
- [ ] convenience functions to build up lists of predicate lists

## Example

Given this HTML page (`htmlText`):

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

The text content of the nodes matching the CSS selector `div.odd p.yes` can be
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

Or easier using the convenience interface `SqueezeSelector`:

```go
doc, _ := html.Parse(strings.NewReader(htmlText))
found := htmlsqueeze.SqueezeSelector(doc, "div.odd p.yes", htmlsqueeze.ExtractChildText)
```
