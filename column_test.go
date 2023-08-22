package carbon

import (
	"bytes"
	"testing"
)

func TestColumn(t *testing.T) {
	// Test that the HTML output of a row component is correct.
	var buf bytes.Buffer
	Column(Raw(`<p>Hello World!</p>`)).Render(&buf)
	if buf.String() != `<div class="bx--col"><p>Hello World!</p></div>` {
		t.Error(buf.String())
	}
}

func TestColumnSm(t *testing.T) {
	// Test that the HTML output of a row component is correct.
	var buf bytes.Buffer
	Column(Raw(`<p>Hello World!</p>`)).Sm(2).Render(&buf)
	if buf.String() != `<div class="bx--col-sm-2"><p>Hello World!</p></div>` {
		t.Error(buf.String())
	}
}
