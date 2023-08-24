package carbon

import (
	"bytes"
	"testing"
)

func TestHtml(t *testing.T) {
	expected := `<!DOCTYPE html><html><head></head><body></body></html>`
	b := &bytes.Buffer{}
	Html().Render(b)
	if b.String() != expected {
		t.Errorf("expected %s, got %s", expected, b.String())
	}
}
