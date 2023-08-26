package carbon

import "io"

type breadcrumb struct {
	attr     []Attr
	children []*breadcrumbItem

	noTrailingSlash bool
}

func Breadcrumb(children ...*breadcrumbItem) *breadcrumb {
	return &breadcrumb{
		attr:     nil,
		children: children,

		noTrailingSlash: false,
	}
}

func (b *breadcrumb) Attr(name string, value string) Component {
	b.attr = append(b.attr, Attr{name, value})
	return b
}

func (b *breadcrumb) NoTrailingSlash(noTrailingSlash bool) *breadcrumb {
	b.noTrailingSlash = noTrailingSlash
	return b
}

func (b *breadcrumb) Render(w io.Writer) {
	w.Write([]byte(`<nav aria-label="Breadcrumb"`))
	renderAttrs(w, b.attr)
	w.Write([]byte(`>`))
	{
		w.Write([]byte(`<ol class="bx--breadcrumb`))
		if b.noTrailingSlash {
			w.Write([]byte(` bx--breadcrumb--no-trailing-slash`))
		}
		w.Write([]byte(`">`))
		{
			for _, child := range b.children {
				child.Render(w)
			}
		}
		w.Write([]byte(`</ol>`))
	}
	w.Write([]byte(`</nav>`))
}
