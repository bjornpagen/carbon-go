package carbon

import "io"

type breadcrumbItem struct {
	children any
	attr     []Attr

	href          string
	isCurrentPage bool
}

func BreadcrumbItem(children ...any) *breadcrumbItem {
	return &breadcrumbItem{
		children:      children,
		attr:          nil,
		href:          "",
		isCurrentPage: false,
	}
}

func (b *breadcrumbItem) Attr(name string, value string) *breadcrumbItem {
	b.attr = append(b.attr, Attr{name, value})
	return b
}

func (b *breadcrumbItem) Href(href string) *breadcrumbItem {
	b.href = href
	return b
}

func (b *breadcrumbItem) IsCurrentPage(isCurrentPage bool) *breadcrumbItem {
	b.isCurrentPage = isCurrentPage
	return b
}

func (b *breadcrumbItem) Render(w io.Writer) {
	ariaCurrent := getAttr(b.attr, "aria-current")
	w.Write([]byte(`<li class="bx--breadcrumb-item`))
	if b.isCurrentPage && ariaCurrent != "page" {
		w.Write([]byte(` bx--breadcrumb-item--current`))
	}
	w.Write([]byte(`">`))
	{
		if b.href != "" {
			Link(b.children).Href(b.href).Attr("aria-current", ariaCurrent).Render(w)
		} else {
			panic("href is required")
		}
	}
	w.Write([]byte(`</li>`))
}
