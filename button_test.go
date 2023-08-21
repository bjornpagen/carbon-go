package carbon

import (
	"bytes"
	"testing"
)

func TestButton(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--primary" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button(Raw("Hello, world!")).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}

	if b == nil {
		t.Error("New() returned nil")
	}
}

func TestButtonDisabled(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--primary bx--btn--disabled" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button(Raw("Hello, world!")).Disabled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonTertiaryHref(t *testing.T) {
	expected := `<a class="bx--btn bx--btn--tertiary" aria-pressed="false" href="https://example.com">Hello, world!</a>`

	b := &bytes.Buffer{}
	Button(Raw("Hello, world!")).Kind("tertiary").Href("https://example.com").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonExpressiveXl(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--expressive bx--btn--xl bx--btn--primary" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button(Raw("Hello, world!")).Expressive(true).Size("xl").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonWithIcon(t *testing.T) {
	svg := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" role="img" style="margin-left: 0" aria-label="Tooltip text" class="bx--btn__icon"><path d="M17 15L17 8 15 8 15 15 8 15 8 17 15 17 15 24 17 24 17 17 24 17 24 15z"></path></svg>`

	expected := `<button class="bx--btn bx--btn--primary" aria-pressed="false">Hello, world!<div aria-hidden="true" class="bx--btn__icon" style="display: contents;" aria-label="Tooltip text">` + svg + `</div></button>`
	b := &bytes.Buffer{}
	Button(Raw("Hello, world!")).Icon(Raw(svg)).IconDescription("Tooltip text").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonIconOnly(t *testing.T) {
	svg := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" role="img" style="margin-left: 0" aria-label="Tooltip text" class="bx--btn__icon"><path d="M17 15L17 8 15 8 15 15 8 15 8 17 15 17 15 24 17 24 17 17 24 17 24 15z"></path></svg>`

	expected := `<button class="bx--btn bx--btn--primary bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y" aria-pressed="false"><span class="bx--assistive-text">Tooltip text</span>` + svg + `</button>`
	b := &bytes.Buffer{}
	Button(nil).Icon(Raw(svg)).IconDescription("Tooltip text").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func BenchmarkButtonRender(b *testing.B) {
	btn := Button(Raw("Hello, world!"))
	for i := 0; i < b.N; i++ {
		btn.Render(&bytes.Buffer{})
	}
}
