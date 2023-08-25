package carbon

import (
	"bytes"
	"testing"
)

func TestLink(t *testing.T) {
	expected := `<a class="bx--link" href="https://www.carbondesignsystem.com/">Carbon Design System</a>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestLinkInline(t *testing.T) {
	expected := `<a class="bx--link bx--link--inline" href="https://www.carbondesignsystem.com/">Carbon Design System</a>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Inline(true).Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestLinkIcon(t *testing.T) {
	expected := `<a class="bx--link" href="https://www.carbondesignsystem.com/">Carbon Design System<div class="bx--link__icon"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true"><path d="M12 12H14V24H12zM18 12H20V24H18z"></path><path d="M4 6V8H6V28a2 2 0 002 2H24a2 2 0 002-2V8h2V6zM8 28V8H24V28zM12 2H20V4H12z"></path></svg></div></a>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Icon(Delete()).Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestLinkLarge(t *testing.T) {
	expected := `<a class="bx--link bx--link--lg" href="https://www.carbondesignsystem.com/">Carbon Design System</a>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Size("lg").Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestLinkSmall(t *testing.T) {
	expected := `<a class="bx--link bx--link--sm" href="https://www.carbondesignsystem.com/">Carbon Design System</a>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Size("sm").Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestLinkDisabled(t *testing.T) {
	expected := `<p class="bx--link bx--link--disabled">Carbon Design System</p>`
	b := bytes.Buffer{}
	Link("Carbon Design System").Href("https://www.carbondesignsystem.com/").Disabled(true).Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
