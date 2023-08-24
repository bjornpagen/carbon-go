package carbon

import (
	"bytes"
	"testing"
)

func TestColumn(t *testing.T) {
	expected := `<div class="bx--col"><div class="bx--row"><p>Hello World!</p></div></div>`
	// Test that the HTML output of a row component is correct.
	var buf bytes.Buffer
	Column(Row(`<p>Hello World!</p>`)).Render(&buf)
	if buf.String() != expected {
		t.Error(buf.String())
	}
}

func TestColumnSm(t *testing.T) {
	expected := `<div class="bx--col-sm-2"><div class="bx--row"><p>Hello World!</p></div></div>`
	// Test that the HTML output of a row component is correct.
	var buf bytes.Buffer
	Column(Row(`<p>Hello World!</p>`)).Sm(2).Render(&buf)
	if buf.String() != expected {
		t.Error(buf.String())
	}
}
