package carbon

import (
	"bytes"
	"testing"
)

func TestStyle(t *testing.T) {
	expected := ``
	b := &bytes.Buffer{}
	Style().Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
