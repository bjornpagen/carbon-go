package carbon

import "io"

type buttonSet struct {
	attrs    []Attr
	children []*button

	stacked bool
}

var _ Component = (*buttonSet)(nil)

func ButtonSet(children ...*button) *buttonSet {
	return &buttonSet{
		attrs:    nil,
		children: children,

		stacked: false,
	}
}

func (b *buttonSet) Attr(name string, value string) *buttonSet {
	b.attrs = append(b.attrs, Attr{name, value})
	return b
}

func (b *buttonSet) Stacked(stacked bool) *buttonSet {
	b.stacked = stacked
	return b
}

func (b *buttonSet) Render(w io.Writer) {
	w.Write([]byte(`<div class="bx--btn-set`))
	if b.stacked {
		w.Write([]byte(` bx--btn-set--stacked`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, b.attrs)
	w.Write([]byte(`>`))
	{
		for _, child := range b.children {
			child.Render(w)
		}
	}
	w.Write([]byte(`</div>`))
}
