package carbon

import (
	"io"
)

type aspectRatio struct {
	children Component
	attrs    []Attr

	ratio string
}

var _ Component = (*aspectRatio)(nil)

func AspectRatio(cs ...Component) *aspectRatio {
	return &aspectRatio{
		children: ternary(len(cs) == 1, cs[0], Fragment(cs...)),
		attrs:    nil,

		ratio: "2x1",
	}
}

func (a *aspectRatio) Attr(name string, value string) Component {
	a.attrs = append(a.attrs, Attr{name, value})
	return a
}

func (a *aspectRatio) Set2x1() *aspectRatio {
	a.ratio = "2x1"
	return a
}

func (a *aspectRatio) Set2x3() *aspectRatio {
	a.ratio = "2x3"
	return a
}

func (a *aspectRatio) Set16x9() *aspectRatio {
	a.ratio = "16x9"
	return a
}

func (a *aspectRatio) Set4x3() *aspectRatio {
	a.ratio = "4x3"
	return a
}

func (a *aspectRatio) Set1x1() *aspectRatio {
	a.ratio = "1x1"
	return a
}

func (a *aspectRatio) Set3x4() *aspectRatio {
	a.ratio = "3x4"
	return a
}

func (a *aspectRatio) Set3x2() *aspectRatio {
	a.ratio = "3x2"
	return a
}

func (a *aspectRatio) Set9x16() *aspectRatio {
	a.ratio = "9x16"
	return a
}

func (a *aspectRatio) Set1x2() *aspectRatio {
	a.ratio = "1x2"
	return a
}

func (a *aspectRatio) Render(w io.Writer) {
	w.Write([]byte(`<div class="bx--aspect-ratio bx--aspect-ratio--`))
	w.Write(yoloBytesUnsafe(a.ratio))
	w.Write([]byte(`"`))
	renderAttrs(w, a.attrs)
	w.Write([]byte(`>`))
	{
		w.Write([]byte(`<div class="bx--aspect-ratio--object">`))
		a.children.Render(w)
		w.Write([]byte(`</div>`))
	}
	w.Write([]byte(`</div>`))
}
