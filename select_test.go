package carbon

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestSelect(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">white</option><option value="g10" class="bx--select-option">g10</option><option value="g80" class="bx--select-option">g80</option><option value="g90" class="bx--select-option">g90</option><option value="g100" class="bx--select-option">g100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white"),
		SelectItem().Value("g10"),
		SelectItem().Value("g80"),
		SelectItem().Value("g90"),
		SelectItem().Value("g100"),
	).LabelText("Carbon theme").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectCustomText(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectInitialValue(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectHelperText(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div><div class="bx--form__helper-text">Carbon offers 4 themes (2 light, 3 dark)</div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).HelperText("Carbon offers 4 themes (2 light, 3 dark)").
		LabelText("Carbon theme").Selected("g10").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectHideLabel(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--visually-hidden">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").HideLabel(true).Selected("g10").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectLightVariant(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select bx--select--light"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").Light(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectInlineVariant(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select bx--select--inline"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input--inline__wrapper"><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").Inline(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectExtraLarge(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input bx--select-input--xl"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").Size("xl").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectSmall(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input bx--select-input--sm"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").Size("sm").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectInvalidState(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select bx--select--invalid"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper" data-invalid><select aria-describedby="error-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" aria-invalid id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__invalid-icon"><path d="M16,2C8.3,2,2,8.3,2,16s6.3,14,14,14s14-6.3,14-14C30,8.3,23.7,2,16,2z M14.9,8h2.2v11h-2.2V8z M16,25 c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22c0.8,0,1.5,0.7,1.5,1.5S16.8,25,16,25z"></path><path fill="none" d="M17.5,23.5c0,0.8-0.7,1.5-1.5,1.5c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22 C16.8,22,17.5,22.7,17.5,23.5z M17.1,8h-2.2v11h2.2V8z" data-icon-path="inner-path" opacity="0"></path></svg></div><div id="error-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--form-requirement">Theme must be a dark variant</div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").
		Invalid(true).InvalidText("Theme must be a dark variant").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectWarningState(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select bx--select--warning"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label">Carbon theme</label><div class="bx--select-input__wrapper"><select id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__invalid-icon bx--select__invalid-icon--warning"><path fill="none" d="M16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Zm-1.125-5h2.25V12h-2.25Z" data-icon-path="inner-path"></path><path d="M16.002,6.1714h-.004L4.6487,27.9966,4.6506,28H27.3494l.0019-.0034ZM14.875,12h2.25v9h-2.25ZM16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Z"></path><path d="M29,30H3a1,1,0,0,1-.8872-1.4614l13-25a1,1,0,0,1,1.7744,0l13,25A1,1,0,0,1,29,30ZM4.6507,28H27.3493l.002-.0033L16.002,6.1714h-.004L4.6487,27.9967Z"></path></svg></div><div id="error-ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--form-requirement">Theme must be a dark variant</div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").
		Warn(true).WarnText("Theme must be a dark variant").Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestSelectDisabled(t *testing.T) {
	rand.Seed(1)
	expected := `<div class="bx--form-item"><div class="bx--select bx--select--disabled"><label for="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" class="bx--label bx--label--disabled">Carbon theme</label><div class="bx--select-input__wrapper"><select disabled id="ccs-XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLS" name="" class="bx--select-input"><option value="white" class="bx--select-option">White</option><option value="g10" selected class="bx--select-option">Gray 10</option><option value="g80" class="bx--select-option">Gray 80</option><option value="g90" class="bx--select-option">Gray 90</option><option value="g100" class="bx--select-option">Gray 100</option></select><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--select__arrow"><path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path></svg></div></div></div>`

	b := &bytes.Buffer{}
	Select(
		SelectItem().Value("white").Text("White"),
		SelectItem().Value("g10").Text("Gray 10"),
		SelectItem().Value("g80").Text("Gray 80"),
		SelectItem().Value("g90").Text("Gray 90"),
		SelectItem().Value("g100").Text("Gray 100"),
	).LabelText("Carbon theme").Selected("g10").
		Disabled(true).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
