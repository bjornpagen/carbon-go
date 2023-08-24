package carbon

import (
	"bytes"
	"testing"
)

func TestOrderedList(t *testing.T) {
	expected := `<ol class="bx--list--ordered"><li class="bx--list__item">Ordered list item</li><li class="bx--list__item">Ordered list item</li><li class="bx--list__item">Ordered list item</li></ol>`
	b := bytes.Buffer{}
	OrderedList(
		ListItem("Ordered list item"),
		ListItem("Ordered list item"),
		ListItem("Ordered list item"),
	).Render(&b)
	if b.String() != expected {
		t.Errorf("Expected\n%s\n\nGot:\n%s", expected, b.String())
	}
}

func TestOrderedListNested(t *testing.T) {
	expected := ``
	b := bytes.Buffer{}
	OrderedList(
		ListItem("Ordered list level 1"),
		OrderedList(
			ListItem("Ordered list level 2"),
			ListItem("Ordered list level 2"),
			OrderedList(
				ListItem("Ordered list level 3"),
				ListItem("Ordered list level 3"),
			).Nested(true),
		).Nested(true),
	).Render(&b)
	if b.String() != expected {
		t.Errorf("Expected\n%s\n\nGot:\n%s", expected, b.String())
	}
}
