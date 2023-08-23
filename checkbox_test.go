package carbon

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestCheckbox(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--checkbox-wrapper"><input type="checkbox" value="" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox"/><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox-label"><span class="bx--checkbox-label-text">Label text</span></label></div>`

	b := &bytes.Buffer{}
	Checkbox().LabelText("Label text").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestCheckboxChecked(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--checkbox-wrapper"><input type="checkbox" value="" checked id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox"/><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox-label"><span class="bx--checkbox-label-text">Label text</span></label></div>`

	b := &bytes.Buffer{}
	Checkbox().LabelText("Label text").Checked(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestCheckboxIndeterminate(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--checkbox-wrapper"><input type="checkbox" value="" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" indeterminate class="bx--checkbox"/><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox-label"><span class="bx--checkbox-label-text">Label text</span></label></div>`

	b := &bytes.Buffer{}
	Checkbox().LabelText("Label text").Indeterminate(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestCheckboxHideLabel(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--checkbox-wrapper"><input type="checkbox" value="" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox"/><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox-label"><span class="bx--checkbox-label-text bx--visually-hidden">Label text</span></label></div>`

	b := &bytes.Buffer{}
	Checkbox().LabelText("Label text").HideLabel(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestCheckboxDisabled(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item bx--checkbox-wrapper"><input type="checkbox" value="" disabled id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox"/><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--checkbox-label"><span class="bx--checkbox-label-text">Label text</span></label></div>`

	b := &bytes.Buffer{}
	Checkbox().LabelText("Label text").Disabled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
