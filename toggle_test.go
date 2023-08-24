package carbon

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestToggle(t *testing.T) {
	rand.Seed(1)
	expected := `<div bx--form-item><input role="switch" type="checkbox" class="bx--toggle-input" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" /><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" aria-label="Push notifications" class="bx--toggle-input__label"><span>Push notifications</span><span class="bx--toggle__switch"><span aria-hidden="true" class="bx--toggle__text--off">Off</span><span aria-hidden="true" class="bx--toggle__text--on">On</span></span></label></div>`

	b := &bytes.Buffer{}
	Toggle().LabelText("Push notifications").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestToggleOn(t *testing.T) {
	rand.Seed(1)
	expected := `<div bx--form-item><input role="switch" type="checkbox" class="bx--toggle-input" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" checked /><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" aria-label="Push notifications" class="bx--toggle-input__label"><span>Push notifications</span><span class="bx--toggle__switch"><span aria-hidden="true" class="bx--toggle__text--off">Off</span><span aria-hidden="true" class="bx--toggle__text--on">On</span></span></label></div>`

	b := &bytes.Buffer{}
	Toggle().LabelText("Push notifications").Toggled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestToggleHiddenLabel(t *testing.T) {
	rand.Seed(1)
	expected := `<div bx--form-item><input role="switch" type="checkbox" class="bx--toggle-input" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" /><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" aria-label="Push notifications" class="bx--toggle-input__label"><span class="bx--visually-hidden">Push notifications</span><span class="bx--toggle__switch"><span aria-hidden="true" class="bx--toggle__text--off">Off</span><span aria-hidden="true" class="bx--toggle__text--on">On</span></span></label></div>`

	b := &bytes.Buffer{}
	Toggle().LabelText("Push notifications").HideLabel(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestToggleDisabled(t *testing.T) {
	rand.Seed(1)
	expected := `<div bx--form-item><input role="switch" type="checkbox" class="bx--toggle-input" id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" disabled /><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" aria-label="Push notifications" class="bx--toggle-input__label"><span>Push notifications</span><span class="bx--toggle__switch"><span aria-hidden="true" class="bx--toggle__text--off">Off</span><span aria-hidden="true" class="bx--toggle__text--on">On</span></span></label></div>`

	b := &bytes.Buffer{}
	Toggle().LabelText("Push notifications").Disabled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
