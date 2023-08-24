package carbon

import "io"

type listItem struct {
	attrs    []Attr
	children any
}

var _ Component = (*listItem)(nil)

func ListItem(children ...any) *listItem {
	return &listItem{
		attrs:    nil,
		children: children,
	}
}

func (l *listItem) Attr(name string, value string) Component {
	l.attrs = append(l.attrs, Attr{name, value})
	return l
}

func (l *listItem) Children(children ...any) *listItem {
	l.children = children
	return l
}

func (l *listItem) Render(w io.Writer) {
	w.Write([]byte(`<li class="bx--list__item"`))
	renderAttrs(w, l.attrs)
	w.Write([]byte(`>`))
	{
		renderAny(w, l.children)
	}
	w.Write([]byte(`</li>`))
}
