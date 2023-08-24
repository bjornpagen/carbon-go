package carbon

import (
	"bytes"
	"testing"
)

func TestGrid(t *testing.T) {
	// Test that the HTML output of a grid component is correct.
	var buf bytes.Buffer
	Grid(Column(Row(`<p>Hello World!</p>`))).Render(&buf)
	if buf.String() != `<div class="bx--grid"><p>Hello World!</p></div>` {
		t.Error(buf.String())
	}
}

func TestGridColumnsRows(t *testing.T) {
	expected := `<div class="bx--grid"><div class="bx--col"><div class="bx--row"><p>Hello World!</p></div></div><div class="bx--col"><div class="bx--row"><p>My name is Jason.</p></div></div></div>`
	// Test that the HTML output of a grid component is correct.
	var buf bytes.Buffer
	Grid(
		Column(Row(`<p>Hello World!</p>`)),
		Column(Row(`<p>My name is Jason.</p>`)),
	).Render(&buf)
	if buf.String() != expected {
		t.Error(buf.String())
	}
}
