package carbon

import "io"

type grid struct {
	content Component
	attrs   []Attr

	as            string
	condensed     bool
	narrow        bool
	fullWidth     bool
	noGutter      bool
	noGutterLeft  bool
	noGutterRight bool
	padding       bool
}

var _ Component = (*grid)(nil)

func Grid(c Component) *grid {
	return &grid{
		content: c,
		attrs:   nil,

		as:            "",
		condensed:     false,
		narrow:        false,
		fullWidth:     false,
		noGutter:      false,
		noGutterLeft:  false,
		noGutterRight: false,
		padding:       false,
	}
}

func (g *grid) Attr(name string, value string) Component {
	g.attrs = append(g.attrs, Attr{name, value})
	return g
}

func (g *grid) As(as string) *grid {
	g.as = as
	return g
}

func (g *grid) Condensed(condensed bool) *grid {
	g.condensed = condensed
	return g
}

func (g *grid) Narrow(narrow bool) *grid {
	g.narrow = narrow
	return g
}

func (g *grid) FullWidth(fullWidth bool) *grid {
	g.fullWidth = fullWidth
	return g
}

func (g *grid) NoGutter(noGutter bool) *grid {
	g.noGutter = noGutter
	return g
}

func (g *grid) NoGutterLeft(noGutterLeft bool) *grid {
	g.noGutterLeft = noGutterLeft
	return g
}

func (g *grid) NoGutterRight(noGutterRight bool) *grid {
	g.noGutterRight = noGutterRight
	return g
}

func (g *grid) Padding(padding bool) *grid {
	g.padding = padding
	return g
}

func (g *grid) Render(w io.Writer) {
	w.Write([]byte("<"))
	if g.as != "" {
		w.Write(yoloBytesUnsafe(g.as))
	} else {
		w.Write([]byte("div"))
	}
	w.Write([]byte(` class="bx--grid`))
	if g.condensed {
		w.Write([]byte(` bx--grid--condensed`))
	}
	if g.narrow {
		w.Write([]byte(` bx--grid--narrow`))
	}
	if g.fullWidth {
		w.Write([]byte(` bx--grid--full-width`))
	}
	if g.noGutter {
		w.Write([]byte(` bx--no-gutter`))
	}
	if g.noGutterLeft {
		w.Write([]byte(` bx--no-gutter--left`))
	}
	if g.noGutterRight {
		w.Write([]byte(` bx--no-gutter--right`))
	}
	if g.padding {
		w.Write([]byte(` bx--grid--padding`))
	}
	w.Write([]byte(`"`))
	renderAttrs(w, g.attrs)
	w.Write([]byte(`>`))
	g.content.Render(w)
	w.Write([]byte("</"))
	if g.as != "" {
		w.Write(yoloBytesUnsafe(g.as))
	} else {
		w.Write([]byte("div"))
	}
	w.Write([]byte(">"))
}
