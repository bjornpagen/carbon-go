package carbon

import (
	"bytes"
	"testing"
)

func TestAspectRatio(t *testing.T) {
	expected := `<div class="bx--aspect-ratio bx--aspect-ratio--2x1"><div class="bx--aspect-ratio--object"></div></div>`

	b := &bytes.Buffer{}
	AspectRatio().Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
