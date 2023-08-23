package carbon

import "io"

type accordionItem struct {
	children []Component
	attrs    []Attr

	title           []Component
	open            bool
	disabled        bool
	iconDescription string
}

var _ Component = (*accordionItem)(nil)

func AccordionItem() *accordionItem {
	return &accordionItem{
		children: nil,
		attrs:    nil,

		title:           []Component{Raw("title")},
		open:            false,
		disabled:        false,
		iconDescription: "Expand/Collapse",
	}
}

func (a *accordionItem) Content(cs ...Component) *accordionItem {
	a.children = cs
	return a
}

func (a *accordionItem) Attr(name string, value string) Component {
	a.attrs = append(a.attrs, Attr{name, value})
	return a
}

func (a *accordionItem) Title(title string) *accordionItem {
	a.title = []Component{Raw(title)}
	return a
}

func (a *accordionItem) TitleComponent(cs ...Component) *accordionItem {
	a.title = cs
	return a
}

func (a *accordionItem) Open(open bool) *accordionItem {
	a.open = open
	return a
}

func (a *accordionItem) Disabled(disabled bool) *accordionItem {
	a.disabled = disabled
	return a
}

func (a *accordionItem) IconDescription(iconDescription string) *accordionItem {
	a.iconDescription = iconDescription
	return a
}

func (a *accordionItem) Render(w io.Writer) {
	w.Write([]byte(`<li class="bx--accordion__item`))
	if a.open {
		w.Write([]byte(` bx--accordion__item--active`))
	}
	if a.disabled {
		w.Write([]byte(` bx--accordion__item--disabled`))
	}
	w.Write([]byte(`">`))
	{
		w.Write([]byte(`<button type="button" class="bx--accordion__heading"`))
		w.Write([]byte(` title="`))
		w.Write(yoloBytesUnsafe(a.iconDescription))
		w.Write([]byte(`"`))
		w.Write([]byte(` aria-expanded="`))
		w.Write(ternary(a.open, []byte(`true`), []byte(`false`)))
		w.Write([]byte(`"`))
		if a.disabled {
			w.Write([]byte(` disabled`))
		}
		w.Write([]byte(`>`))
		{
			ChevronRight().Attr("class", "bx--accordion__arrow").Attr("aria-label", a.iconDescription).Render(w)
			w.Write([]byte(`<div class="bx--accordion__title">`))
			{
				for _, child := range a.title {
					child.Render(w)
				}
			}
			w.Write([]byte(`</div>`))
		}
		renderAttrs(w, a.attrs)
		w.Write([]byte(`</button>`))
		w.Write([]byte(`<div class="bx--accordion__content">`))
		{
			for _, child := range a.children {
				child.Render(w)
			}
		}
		w.Write([]byte(`</div>`))
	}
	w.Write([]byte(`</li>`))
}
