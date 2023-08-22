package carbon

import "io"

type accordionItem struct {
	children []Component
	attrs    []Attr

	title           string
	open            bool
	disabled        bool
	iconDescription string
}

var _ Component = (*accordionItem)(nil)

func AccordionItem(children ...Component) *accordionItem {
	return &accordionItem{
		children: children,
		attrs:    nil,

		title:           "title",
		open:            false,
		disabled:        false,
		iconDescription: "Expand/Collapse",
	}
}

func (a *accordionItem) Attr(name string, value string) Component {
	a.attrs = append(a.attrs, Attr{name, value})
	return a
}

func (a *accordionItem) Title(title string) *accordionItem {
	a.title = title
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
	panic("not implemented")
}
