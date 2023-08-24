package carbon

import "io"

type row struct {
	children any
	attrs    []Attr

	condensed     bool
	narrow        bool
	noGutter      bool
	noGutterLeft  bool
	noGutterRight bool
	padding       bool
}

var _ Component = (*row)(nil)

func Row(a ...any) *row {
	return &row{
		children: a,
		attrs:    nil,

		condensed:     false,
		narrow:        false,
		noGutter:      false,
		noGutterLeft:  false,
		noGutterRight: false,
		padding:       false,
	}
}

func (r *row) Attr(name string, value string) Component {
	r.attrs = append(r.attrs, Attr{name, value})
	return r
}

func (r *row) Condensed(condensed bool) *row {
	r.condensed = condensed
	return r
}

func (r *row) Narrow(narrow bool) *row {
	r.narrow = narrow
	return r
}

func (r *row) NoGutter(noGutter bool) *row {
	r.noGutter = noGutter
	return r
}

func (r *row) NoGutterLeft(noGutterLeft bool) *row {
	r.noGutterLeft = noGutterLeft
	return r
}

func (r *row) NoGutterRight(noGutterRight bool) *row {
	r.noGutterRight = noGutterRight
	return r
}

func (r *row) Padding(padding bool) *row {
	r.padding = padding
	return r
}

func (r *row) Render(w io.Writer) {
	w.Write([]byte(`<div class="bx--row`))
	if r.condensed {
		w.Write([]byte(` bx--row--condensed`))
	}
	if r.narrow {
		w.Write([]byte(` bx--row--narrow`))
	}
	if r.noGutter {
		w.Write([]byte(` bx--no-gutter`))
	}
	if r.noGutterLeft {
		w.Write([]byte(` bx--no-gutter--left`))
	}
	if r.noGutterRight {
		w.Write([]byte(` bx--no-gutter--right`))
	}
	if r.padding {
		w.Write([]byte(` bx--row-padding`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, r.attrs)
	w.Write([]byte(`>`))
	{
		renderAny(w, r.children)
	}
	w.Write([]byte("</div>"))
}
