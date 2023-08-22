package carbon

import (
	"bytes"
	"io"
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

func TestGridRowsColumns(t *testing.T) {
	expected := `<div class="bx--grid"><div class="bx--row"><div class="bx--col"><p>Hello World!</p></div></div><div class="bx--row"><div class="bx--col"><p>My name is Jason.</p></div></div></div>`
	// Test that the HTML output of a grid component is correct.
	var buf bytes.Buffer
	Grid(
		Row(
			Column(Raw(`<p>Hello World!</p>`)),
		),
		Row(
			Column(Raw(`<p>My name is Jason.</p>`)),
		),
	).Render(&buf)
	if buf.String() != expected {
		t.Error(buf.String())
	}
}

func BenchmarkGridRowsColumns(b *testing.B) {
	g := Grid(
		Row(
			Column(Raw(`<p>Hello World!</p>`)),
		),
		Row(
			Column(Raw(`<p>My name is Jason.</p>`)),
		),
	)
	for i := 0; i < b.N; i++ {
		g.Render(io.Discard)
	}
}
