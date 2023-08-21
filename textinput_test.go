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
	expected := `<div class="bx--form-item bx--text-input-wrapper bx--text-input-wrapper--readonly"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet width="16" height="16" class="bx--text-input__readonly-icon"><path d="M30 28.6L3.4 2 2 3.4l10.1 10.1L4 21.6V28h6.4l8.1-8.1L28.6 30 30 28.6zM9.6 26H6v-3.6l7.5-7.5 3.6 3.6L9.6 26zM29.4 6.2L29.4 6.2l-3.6-3.6c-.8-.8-2-.8-2.8 0l0 0 0 0-8 8 1.4 1.4L20 8.4l3.6 3.6L20 15.6l1.4 1.4 8-8C30.2 8.2 30.2 7 29.4 6.2L29.4 6.2zM25 10.6L21.4 7l3-3L28 7.6 25 10.6z"></path></svg><input id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." readonly class="bx--text-input"/></div></div></div>`

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
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div data-invalid class="bx--text-input__field-wrapper"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet width="16" height="16" class="bx--text-input__invalid-icon"><path d="M16,2C8.3,2,2,8.3,2,16s6.3,14,14,14s14-6.3,14-14C30,8.3,23.7,2,16,2z M14.9,8h2.2v11h-2.2V8z M16,25 c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22c0.8,0,1.5,0.7,1.5,1.5S16.8,25,16,25z"></path><path fill="none" d="M17.5,23.5c0,0.8-0.7,1.5-1.5,1.5c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22 C16.8,22,17.5,22.7,17.5,23.5z M17.1,8h-2.2v11h2.2V8z" data-icon-path="inner-path" opacity="0"></path></svg><input data-invalid aria-invalid aria-describedby="error-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input bx--text-input--invalid"/></div><div class="bx--form-requirement" id="error-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS"></div></div></div>`

	b := &bytes.Buffer{}
	TextInput().Invalid(true).LabelText("User name").Placeholder("Enter user name...").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestTextInputWarningState(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--text-input-wrapper"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">User name</label><div class="bx--text-input__field-outer-wrapper"><div data-warn class="bx--text-input__field-wrapper bx--text-input__field-wrapper--warning"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet width="16" height="16" class="bx--text-input__invalid-icon bx--text-input__invalid-icon--warning"><path fill="none" d="M16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Zm-1.125-5h2.25V12h-2.25Z" data-icon-path="inner-path"></path><path d="M16.002,6.1714h-.004L4.6487,27.9966,4.6506,28H27.3494l.0019-.0034ZM14.875,12h2.25v9h-2.25ZM16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Z"></path><path d="M29,30H3a1,1,0,0,1-.8872-1.4614l13-25a1,1,0,0,1,1.7744,0l13,25A1,1,0,0,1,29,30ZM4.6507,28H27.3493l.002-.0033L16.002,6.1714h-.004L4.6487,27.9967Z"></path></svg><input data-warn aria-describedby="warn-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" placeholder="Enter user name..." class="bx--text-input bx--text-input--warning"/></div><div class="bx--form-requirement" id="warn-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS"></div></div></div>`

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

func BenchmarkTextInputRender(b *testing.B) {
	input := TextInput()
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		input.Render(&buf)
	}
}
