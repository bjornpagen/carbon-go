package carbon

import (
	"io"
	"slices"
)

type contentSwitcher struct {
	children []*switch_
	attrs    []Attr

	selectedIndex int8
	size          string
}

var _ Component = (*contentSwitcher)(nil)

func ContentSwitcher(cs ...*switch_) *contentSwitcher {
	return &contentSwitcher{
		children: cs,
		attrs:    nil,

		selectedIndex: 0,
		size:          "",
	}
}

func (c *contentSwitcher) Attr(name string, value string) *contentSwitcher {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *contentSwitcher) SelectedIndex(selectedIndex int8) *contentSwitcher {
	c.selectedIndex = selectedIndex
	return c
}

func (c *contentSwitcher) Size(size string) *contentSwitcher {
	var sizes = []string{"", "sm", "xl"}
	if !slices.Contains(sizes, size) {
		panic("Size() received an invalid value")
	}
	c.size = size
	return c
}

func (c *contentSwitcher) Render(w io.Writer) {
	w.Write([]byte(`<div role="tablist" class="bx--content-switcher`))
	if c.size != "" {
		w.Write([]byte(` bx--content-switcher--`))
		io.WriteString(w, c.size)
	}
	w.Write([]byte(`">`))
	{
		for i, child := range c.children {
			if c.selectedIndex == int8(i) {
				child.selected = true
			} else {
				child.selected = false
			}
			child.Render(w)
		}
	}
	w.Write([]byte(`</div>`))
}
