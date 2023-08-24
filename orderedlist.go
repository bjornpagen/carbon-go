package carbon

import "io"

type orderedList struct {
	attrs    []Attr
	children []any

	nested     bool
	native     bool
	expressive bool
}

var _ Component = (*orderedList)(nil)

func OrderedList(children ...any) *orderedList {
	return &orderedList{
		attrs:    nil,
		children: children,

		nested:     false,
		native:     false,
		expressive: false,
	}
}

func (o *orderedList) Attr(name string, value string) Component {
	o.attrs = append(o.attrs, Attr{name, value})
	return o
}

func (o *orderedList) Children(children ...any) *orderedList {
	o.children = children
	return o
}

func (o *orderedList) Nested(nested bool) *orderedList {
	o.nested = nested
	return o
}

func (o *orderedList) Native(native bool) *orderedList {
	o.native = native
	return o
}

func (o *orderedList) Expressive(expressive bool) *orderedList {
	o.expressive = expressive
	return o
}

func (o *orderedList) Render(w io.Writer) {
	w.Write([]byte(`<ol class="bx--list--ordered`))
	if o.native {
		w.Write([]byte(`--native`))
	}
	if o.nested {
		w.Write([]byte(` bx--list--nested`))
	}
	if o.expressive {
		w.Write([]byte(` bx--list--expressive`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, o.attrs)
	w.Write([]byte(`>`))
	{
		renderAny(w, o.children)
	}
	w.Write([]byte(`</ol>`))
}
