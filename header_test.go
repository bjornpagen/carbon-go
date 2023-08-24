package carbon

import (
	"bytes"
	"testing"
)

func TestHeader(t *testing.T) {
	expected := `<header class="bx--header" aria-label=""><a class="bx--header__name" href=""><span class="bx--header__name--prefix">IBM</span></a></header>`
	b := bytes.Buffer{}
	Header().Company("IBM").Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
