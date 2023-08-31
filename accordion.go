package carbon

import (
	"io"
	"slices"
)

type accordion struct {
	children []*accordionItem
	attrs    []Attr

	align    string
	size     string
	disabled bool
}

var _ Component = (*accordion)(nil)

func Accordion(cs ...*accordionItem) *accordion {
	return &accordion{
		children: cs,
		attrs:    nil,

		align:    "end",
		size:     "",
		disabled: false,
	}
}

func (a *accordion) Attr(name string, value string) *accordion {
	a.attrs = append(a.attrs, Attr{name, value})
	return a
}

func (a *accordion) Align(align string) *accordion {
	alignments := []string{"start", "end"}
	if !slices.Contains(alignments, align) {
		panic("Align() received an invalid value")
	}
	a.align = align
	return a
}

func (a *accordion) Size(size string) *accordion {
	sizes := []string{"", "sm", "lg"}
	if !slices.Contains(sizes, size) {
		panic("Size() received an invalid value")
	}
	a.size = size
	return a
}

func (a *accordion) Disabled(disabled bool) *accordion {
	a.disabled = disabled
	return a
}

func (a *accordion) Render(w io.Writer) {
	w.Write([]byte(`<ul class="bx--accordion`))
	switch a.align {
	case "start":
		w.Write([]byte(` bx--accordion--start`))
	case "end":
		w.Write([]byte(` bx--accordion--end`))
	}
	switch a.size {
	case "sm":
		w.Write([]byte(` bx--accordion--sm`))
	case "lg":
		w.Write([]byte(` bx--accordion--lg`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, a.attrs)
	w.Write([]byte(`>`))
	for _, child := range a.children {
		child.Render(w)
	}
	w.Write([]byte(`</ul>`))
}
