package carbon

import "io"

type content struct {
	children []Component
	attrs    []Attr

	id string
}

var _ Component = (*content)(nil)

func Content(cs ...Component) *content {
	return &content{
		children: cs,
		attrs:    nil,

		id: "main-content",
	}
}

func (c *content) Attr(name string, value string) Component {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *content) Id(id string) *content {
	c.id = id
	return c
}

func (c *content) Render(w io.Writer) {
	w.Write([]byte(`<main id="`))
	w.Write(yoloBytesUnsafe(c.id))
	w.Write([]byte(`" class="bx--content">`))
	for _, child := range c.children {
		child.Render(w)
	}
	w.Write([]byte(`</main>`))
}
