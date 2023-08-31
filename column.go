package carbon

import (
	"io"
	"strconv"

	"golang.org/x/exp/slices"
)

type column struct {
	children []*row
	attrs    []Attr

	noGutter      bool
	noGutterLeft  bool
	noGutterRight bool
	padding       bool
	aspectRatio   string

	sm  int8
	md  int8
	lg  int8
	xlg int8
	max int8
}

var _ Component = (*column)(nil)

func Column(c ...*row) *column {
	return &column{
		children: c,
		attrs:    nil,

		noGutter:      false,
		noGutterLeft:  false,
		noGutterRight: false,
		padding:       false,
		aspectRatio:   "",

		sm:  0,
		md:  0,
		lg:  0,
		xlg: 0,
		max: 0,
	}
}

func (c *column) Attr(name string, value string) *column {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *column) NoGutter(noGutter bool) *column {
	c.noGutter = noGutter
	return c
}

func (c *column) NoGutterLeft(noGutterLeft bool) *column {
	c.noGutterLeft = noGutterLeft
	return c
}

func (c *column) NoGutterRight(noGutterRight bool) *column {
	c.noGutterRight = noGutterRight
	return c
}

func (c *column) Padding(padding bool) *column {
	c.padding = padding
	return c
}

func (c *column) AspectRatio(aspectRatio string) *column {
	allowedAspectRatios := []string{"2x1", "16x9", "9x16", "1x2", "4x3", "3x4", "1x1"}
	if slices.Contains(allowedAspectRatios, aspectRatio) == false {
		panic("invalid aspect ratio")
	}
	c.aspectRatio = aspectRatio
	return c
}

func (c *column) Sm(b int8) *column {
	if b < 1 || b > 4 {
		panic("invalid breakpoint")
	}
	c.sm = b
	return c
}

func (c *column) Md(b int8) *column {
	if b < 1 || b > 8 {
		panic("invalid breakpoint")
	}
	c.md = b
	return c
}

func (c *column) Lg(b int8) *column {
	if b < 1 || b > 16 {
		panic("invalid breakpoint")
	}
	c.lg = b
	return c
}

func (c *column) Xl(b int8) *column {
	if b < 1 || b > 16 {
		panic("invalid breakpoint")
	}
	c.xlg = b
	return c
}

func (c *column) Max(b int8) *column {
	if b < 1 || b > 16 {
		panic("invalid breakpoint")
	}
	c.max = b
	return c
}

func (c *column) Render(w io.Writer) {
	var _isColumnClass bool

	w.Write([]byte(`<div class="`))

	var allowedBreakpoints = []string{"sm", "md", "lg", "xlg", "max"}
	for _, breakpoint := range allowedBreakpoints {
		var v int8
		switch breakpoint {
		case "sm":
			v = c.sm
		case "md":
			v = c.md
		case "lg":
			v = c.lg
		case "xlg":
			v = c.xlg
		case "max":
			v = c.max
		}

		if v != 0 {
			if _isColumnClass {
				w.Write([]byte(" "))
			}
			w.Write([]byte("bx--col-"))
			io.WriteString(w, breakpoint)
			w.Write([]byte("-"))
			io.WriteString(w, strconv.Itoa(int(v)))
			_isColumnClass = true
		}
	}

	if _isColumnClass == false {
		w.Write([]byte("bx--col"))
	}
	if c.noGutter {
		w.Write([]byte(" bx--no-gutter"))
	}
	if c.noGutterLeft {
		w.Write([]byte(" bx--no-gutter--left"))
	}
	if c.noGutterRight {
		w.Write([]byte(" bx--no-gutter--right"))
	}
	if c.aspectRatio != "" {
		w.Write([]byte(" bx--aspect-ratio bx--aspect-ratio--"))
		io.WriteString(w, c.aspectRatio)
	}
	if c.padding {
		w.Write([]byte(" bx--col-padding"))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, c.attrs)
	w.Write([]byte(`>`))
	{
		for _, child := range c.children {
			child.Render(w)
		}
	}
	w.Write([]byte("</div>"))
}
