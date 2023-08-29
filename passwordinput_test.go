package carbon

import (
	"bytes"
	"testing"
)

func TestPasswordInput(t *testing.T) {
	expected := `<div class="bx--form-item bx--text-input-wrapper bx--password-input-wrapper"><label for="ccs-epiAgkSgKyfXohPXgmzuvTDcWtbTshaFrfUY" class="bx--label">Password</label><div class="bx--text-input__field-outer-wrapper"><div class="bx--text-input__field-wrapper"><input id="ccs-epiAgkSgKyfXohPXgmzuvTDcWtbTshaFrfUY" placeholder="Enter your password" type="password" class="bx--text-input bx--password-input" /><button type="button" disabled class="bx--text-input--password__visibility__toggle bx--btn bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y bx--tooltip--bottom bx--tooltip--align-center"><span class="bx--assistive-text">Show password</span><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" aria-hidden="true" class="bx--icon-visibility-on"><path d="M30.94,15.66A16.69,16.69,0,0,0,16,5,16.69,16.69,0,0,0,1.06,15.66a1,1,0,0,0,0,.68A16.69,16.69,0,0,0,16,27,16.69,16.69,0,0,0,30.94,16.34,1,1,0,0,0,30.94,15.66ZM16,25c-5.3,0-10.9-3.93-12.93-9C5.1,10.93,10.7,7,16,7s10.9,3.93,12.93,9C26.9,21.07,21.3,25,16,25Z"></path><path d="M16,10a6,6,0,1,0,6,6A6,6,0,0,0,16,10Zm0,10a4,4,0,1,1,4-4A4,4,0,0,1,16,20Z"></path></svg></button></div></div></div>`
	b := bytes.Buffer{}
	PasswordInput("Password").Placeholder("Enter your password").Render(&b)
	if b.String() != expected {
		t.Errorf("Expected\n%s\n, got:\n%s", expected, b.String())
	}
}
