package carbon

import (
	"bytes"
	"io"
	"testing"
)

func TestButtonPrimary(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--primary" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}

	if b == nil {
		t.Error("New() returned nil")
	}
}

func TestButtonSecondary(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--secondary" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Kind("secondary").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonTertiary(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--tertiary" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Kind("tertiary").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonGhost(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--ghost" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Kind("ghost").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonDanger(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--danger" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Kind("danger").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonDangerTertiaryIconOnlyNoDescription(t *testing.T) {
	defer func() { _ = recover() }()

	b := &bytes.Buffer{}
	Button().Kind("danger-tertiary").Icon(Delete()).Render(b)

	t.Error("expected panic")
}

func TestButtonDangerTertiaryIconOnly(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--danger-tertiary bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y bx--btn--icon-only--bottom bx--tooltip--align-center" aria-pressed="false"><span class="bx--assistive-text">Delete</span><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--btn__icon" aria-hidden="true" aria-label="Delete"><path d="M12 12H14V24H12zM18 12H20V24H18z"></path><path d="M4 6V8H6V28a2 2 0 002 2H24a2 2 0 002-2V8h2V6zM8 28V8H24V28zM12 2H20V4H12z"></path></svg></button>`

	b := &bytes.Buffer{}
	Button().Kind("danger-tertiary").Icon(Delete()).IconDescription("Delete").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonDisabled(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--primary bx--btn--disabled" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Disabled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonDangerGhost(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--danger-ghost" aria-pressed="false">Hello, world!</button>`

	b := &bytes.Buffer{}
	Button("Hello, world!").Kind("danger-ghost").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonIconOnly(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--primary bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y bx--btn--icon-only--bottom bx--tooltip--align-center" aria-pressed="false"><span class="bx--assistive-text">Add</span><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--btn__icon" aria-hidden="true" aria-label="Add"><path d="M16 16V8h-2v8H6v2h8v8h2v-8h8v-2z"></path></svg></button>`

	b := &bytes.Buffer{}
	Button().Kind("primary").Icon(Add()).IconDescription("Add").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonIconOnlyLink(t *testing.T) {
	expected := `<a class="bx--btn bx--btn--primary bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y bx--btn--icon-only--bottom bx--tooltip--align-center" aria-pressed="false" href="#"><span class="bx--assistive-text">Add</span><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--btn__icon" aria-hidden="true" aria-label="Add"><path d="M16 16V8h-2v8H6v2h8v8h2v-8h8v-2z"></path></svg></a>`

	b := &bytes.Buffer{}
	Button().Kind("primary").Icon(Add()).IconDescription("Add").Href("#").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestButtonIconOnlyGhostSelected(t *testing.T) {
	expected := `<button class="bx--btn bx--btn--ghost bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y bx--btn--icon-only--bottom bx--tooltip--align-center bx--btn--selected" aria-pressed="true"><span class="bx--assistive-text">Add</span><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--btn__icon" aria-hidden="true" aria-label="Add"><path d="M16 16V8h-2v8H6v2h8v8h2v-8h8v-2z"></path></svg></button>`

	b := &bytes.Buffer{}
	Button().Kind("ghost").Icon(Add()).IconDescription("Add").Selected(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func BenchmarkButtonRender(b *testing.B) {
	btn := Button("Hello, world!")
	for i := 0; i < b.N; i++ {
		btn.Render(io.Discard)
	}
}
