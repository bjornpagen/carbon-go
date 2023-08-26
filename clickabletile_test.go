package carbon

import (
	"bytes"
	"testing"
)

func TestClickableTile(t *testing.T) {
	expected := `<a class="bx--link bx--tile bx--tile--clickable" href="https://www.carbondesignsystem.com/">Carbon Design System</a>`
	b := bytes.Buffer{}
	ClickableTile("Carbon Design System").Href("https://www.carbondesignsystem.com/").Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestClickableTileDisabled(t *testing.T) {
	expected := `<p class="bx--link bx--link--disabled bx--tile bx--tile--clickable">Carbon Design System</p>`
	b := bytes.Buffer{}
	ClickableTile("Carbon Design System").Href("https://www.carbondesignsystem.com/").Disabled(true).Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
