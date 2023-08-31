package carbon

import (
	"io"
)

type section struct {
	children any
	attrs    []Attr
}

var _ Component = (*section)(nil)

func Section(a ...any) *section {
	return &section{
		children: a,
		attrs:    nil,
	}
}

func (s *section) Attr(name string, value string) *section {
	s.attrs = append(s.attrs, Attr{name, value})
	return s
}

func (s *section) Render(w io.Writer) {
	w.Write([]byte(`<section`))
	renderAttrs(w, s.attrs)
	w.Write([]byte(`>`))
	renderAny(w, s.children)
	w.Write([]byte(`</section>`))
}
