package carbon

import (
	"bytes"
	"testing"
)

func TestContentSwitcher(t *testing.T) {
	expected := `<div role="tablist" class="bx--content-switcher"><button type="button" role="tab" tabindex="0" aria-selected class="bx--content-switcher-btn bx--content-switcher--selected"><span class="bx--content-switcher__label">Latest news</span></button><button type="button" role="tab" tabindex="-1" class="bx--content-switcher-btn"><span class="bx--content-switcher__label">Trending</span></button></div>`

	b := &bytes.Buffer{}
	ContentSwitcher(
		Switch("Latest news"),
		Switch("Trending"),
	).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestContentInitialSelectedIndex(t *testing.T) {
	expected := `<div role="tablist" class="bx--content-switcher"><button type="button" role="tab" tabindex="-1" class="bx--content-switcher-btn"><span class="bx--content-switcher__label">Latest news</span></button><button type="button" role="tab" tabindex="0" aria-selected class="bx--content-switcher-btn bx--content-switcher--selected"><span class="bx--content-switcher__label">Trending</span></button><button type="button" role="tab" tabindex="-1" class="bx--content-switcher-btn"><span class="bx--content-switcher__label">Recommended</span></button></div>`

	b := &bytes.Buffer{}
	ContentSwitcher(
		Switch("Latest news"),
		Switch("Trending"),
		Switch("Recommended"),
	).SelectedIndex(1).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestContentSizeXL(t *testing.T) {
	expected := `<div role="tablist" class="bx--content-switcher bx--content-switcher--xl"><button type="button" role="tab" tabindex="0" aria-selected class="bx--content-switcher-btn bx--content-switcher--selected"><span class="bx--content-switcher__label">Latest news</span></button><button type="button" role="tab" tabindex="-1" class="bx--content-switcher-btn"><span class="bx--content-switcher__label">Trending</span></button></div>`

	b := &bytes.Buffer{}
	ContentSwitcher(
		Switch("Latest news"),
		Switch("Trending"),
	).Size("xl").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestContentSwitcherDisabled(t *testing.T) {
	expected := `<div role="tablist" class="bx--content-switcher"><button type="button" role="tab" tabindex="0" aria-selected disabled class="bx--content-switcher-btn bx--content-switcher--selected"><span class="bx--content-switcher__label">Latest news</span></button><button type="button" role="tab" tabindex="-1" disabled class="bx--content-switcher-btn"><span class="bx--content-switcher__label">Trending</span></button></div>`

	b := &bytes.Buffer{}
	ContentSwitcher(
		Switch("Latest news").Disabled(true),
		Switch("Trending").Disabled(true),
	).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
