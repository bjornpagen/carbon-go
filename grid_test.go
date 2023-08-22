package carbon

import (
	"bytes"
	"testing"
)

func TestGrid(t *testing.T) {
	// Test that the HTML output of a grid component is correct.
	var buf bytes.Buffer
	Grid(Raw(`<p>Hello World!</p>`)).Render(&buf)
	if buf.String() != `<div class="bx--grid"><p>Hello World!</p></div>` {
		t.Error(buf.String())
	}
}
