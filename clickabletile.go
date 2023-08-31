package carbon

import (
	"io"

	"github.com/valyala/bytebufferpool"
)

type clickableTile struct {
	children any
	attrs    []Attr

	clicked  bool
	light    bool
	disabled bool
	href     string
}

func ClickableTile(children ...any) *clickableTile {
	return &clickableTile{
		children: children,
		attrs:    nil,

		clicked:  false,
		light:    false,
		disabled: false,
		href:     "",
	}
}

func (c *clickableTile) Attr(name string, value string) *clickableTile {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *clickableTile) Clicked(clicked bool) *clickableTile {
	c.clicked = clicked
	return c
}

func (c *clickableTile) Light(light bool) *clickableTile {
	c.light = light
	return c
}

func (c *clickableTile) Disabled(disabled bool) *clickableTile {
	c.disabled = disabled
	return c
}

func (c *clickableTile) Href(href string) *clickableTile {
	c.href = href
	return c
}

func (c *clickableTile) Render(w io.Writer) {
	classBuilder := bytebufferpool.Get()
	defer bytebufferpool.Put(classBuilder)

	classBuilder.WriteString("bx--tile bx--tile--clickable")
	if c.clicked {
		classBuilder.WriteString(" bx--tile--is-clicked")
	}
	if c.light {
		classBuilder.WriteString(" bx--tile--light")
	}
	classBuilder.WriteString(getAttr(c.attrs, "class"))

	Link(c.children).Disabled(c.disabled).Href(c.href).Attr("class", classBuilder.String()).Render(w)
}
