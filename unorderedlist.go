package carbon

import "io"

type unorderedList struct {
	attrs    []Attr
	children any

	nested     bool
	expressive bool
}

var _ Component = (*unorderedList)(nil)

func UnorderedList(children ...any) *unorderedList {
	return &unorderedList{
		attrs:    nil,
		children: children,

		nested:     false,
		expressive: false,
	}
}

func (u *unorderedList) Attr(name string, value string) *unorderedList {
	u.attrs = append(u.attrs, Attr{name, value})
	return u
}

func (u *unorderedList) Nested(nested bool) *unorderedList {
	u.nested = nested
	return u
}

func (u *unorderedList) Expressive(expressive bool) *unorderedList {
	u.expressive = expressive
	return u
}

func (u *unorderedList) Render(w io.Writer) {
	w.Write([]byte(`<ul class="bx--list--unordered`))
	if u.nested {
		w.Write([]byte(` bx--list--nested`))
	}
	if u.expressive {
		w.Write([]byte(` bx--list--expressive`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, u.attrs)
	w.Write([]byte(`>`))
	{
		renderAny(w, u.children)
	}
	w.Write([]byte(`</ul>`))
}
