package carbon

import (
	"bytes"
	"io"
	"testing"
)

func TestRow(t *testing.T) {
	// Test that the HTML output of a row component is correct.
	var buf bytes.Buffer
	Row(Raw(`<p>Hello World!</p>`)).Render(&buf)
	if buf.String() != `<div class="bx--row"><p>Hello World!</p></div>` {
		t.Error(buf.String())
	}
}

func BenchmarkRow(b *testing.B) {
	r := Row(Raw(`<p>Hello World!</p>`))
	for i := 0; i < b.N; i++ {
		r.Render(io.Discard)
	}
}
