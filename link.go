package carbon

import (
	"io"
	"slices"
)

type link struct {
	attr     []Attr
	children any

	size     string
	href     string
	inline   bool
	icon     Component
	disabled bool
	visited  bool
}

func Link(children ...any) *link {
	return &link{
		attr:     nil,
		children: children,

		size:     "",
		href:     "",
		inline:   false,
		icon:     nil,
		disabled: false,
		visited:  false,
	}
}

func (l *link) Attr(name string, value string) *link {
	l.attr = append(l.attr, Attr{name, value})
	return l
}

func (l *link) Size(size string) *link {
	sizes := []string{"", "sm", "lg"}
	if !slices.Contains(sizes, size) {
		panic("invalid size")
	}
	l.size = size
	return l
}

func (l *link) Href(href string) *link {
	l.href = href
	return l
}

func (l *link) Inline(inline bool) *link {
	l.inline = inline
	return l
}

func (l *link) Icon(icon Component) *link {
	l.icon = icon
	return l
}

func (l *link) Disabled(disabled bool) *link {
	l.disabled = disabled
	return l
}

func (l *link) Visited(visited bool) *link {
	l.visited = visited
	return l
}

func (l *link) Render(w io.Writer) {
	if l.disabled {
		w.Write([]byte(`<p class="bx--link bx--link--disabled`))
		if l.inline {
			w.Write([]byte(` bx--link--inline`))
		}
		if l.visited {
			w.Write([]byte(` bx--link--visited`))
		}
		w.Write([]byte(` `))
		io.WriteString(w, getAttr(l.attr, "class"))
		w.Write([]byte(`"`))
		renderAttrsWithoutClass(w, l.attr)
		w.Write([]byte(`>`))
		{
			renderAny(w, l.children)
			if !l.inline && l.icon != nil {
				w.Write([]byte(`<div class="bx--link__icon">`))
				l.icon.Render(w)
				w.Write([]byte(`</div>`))
			}
		}
		w.Write([]byte(`</p>`))
		return
	}

	w.Write([]byte(`<a class="bx--link`))
	if l.inline {
		w.Write([]byte(` bx--link--inline`))
	}
	if l.visited {
		w.Write([]byte(` bx--link--visited`))
	}
	if l.size != "" {
		w.Write([]byte(` bx--link--`))
		io.WriteString(w, l.size)
	}
	w.Write([]byte(` `))
	io.WriteString(w, getAttr(l.attr, "class"))
	w.Write([]byte(`"`))
	if l.href != "" {
		w.Write([]byte(` href="`))
		io.WriteString(w, l.href)
		w.Write([]byte(`"`))
	}
	renderAttrsWithoutClass(w, l.attr)
	w.Write([]byte(`>`))
	{
		renderAny(w, l.children)
		if !l.inline && l.icon != nil {
			w.Write([]byte(`<div class="bx--link__icon">`))
			l.icon.Render(w)
			w.Write([]byte(`</div>`))
		}
	}
	w.Write([]byte(`</a>`))
}
