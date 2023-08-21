package carbon

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestTextInputDefault(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputWithHelperText(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input aria-describedby="helper-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input"/></div><div id="helper-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--form__helper-text">Your user name is associated with your email</div></div></div>`

	b := &bytes.Buffer{}
	TextInput().LabelText("User name").HelperText("Your user name is associated with your email").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputHiddenLabel(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--visually-hidden">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().HideLabel(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputLightVariant(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper bx--text-input-wrapper--light"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input bx--text-input--light"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Light(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputInlineVariant(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper bx--text-input-wrapper--inline"><div class="bx--text-input__label-helper-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--label--inline">User name</label></div><div class="bx--text-input__field-outer-wrapper bx--text-input__field-outer-wrapper--inline"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Inline(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputReadonlyVariant(t *testing.T) {
	rand.Seed(1)
	expected := ``

	b := &bytes.Buffer{}
	TextInput().Readonly(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputXLSize(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--label--inline--xl">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input bx--text-input--xl"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Size("xl").LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputSmSize(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--label--inline--sm">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input bx--text-input--sm"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Size("sm").LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputInvalidState(t *testing.T) {
	rand.Seed(1)
	expected := ``

	b := &bytes.Buffer{}
	TextInput().Invalid(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputWarningState(t *testing.T) {
	rand.Seed(1)
	expected := ``

	b := &bytes.Buffer{}
	TextInput().Warn(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputDisabled(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--label--disabled">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input disabled id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Disabled(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputAttr(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input" hx-swap="outerHTML"/></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().LabelText("User name").Placeholder("Enter user name...").Attr("hx-swap", "outerHTML").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func BenchmarkTextInput(b *testing.B) {
	input := TextInput()
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		input.Render(&buf)
	}
}
