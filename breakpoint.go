package carbon

import "io"

type breakpoint struct {
	attr []Attr

	sm  any
	md  any
	lg  any
	xlg any
	max any
}

func Breakpoint(sm ...any) *breakpoint {
	return &breakpoint{
		attr: nil,

		sm:  sm,
		md:  nil,
		lg:  nil,
		xlg: nil,
		max: nil,
	}
}

func (b *breakpoint) Attr(name string, value string) Component {
	b.attr = append(b.attr, Attr{name, value})
	return b
}

func (b *breakpoint) Md(md ...any) *breakpoint {
	b.md = md
	return b
}

func (b *breakpoint) Lg(lg ...any) *breakpoint {
	b.lg = lg
	return b
}

func (b *breakpoint) Xlg(xlg ...any) *breakpoint {
	b.xlg = xlg
	return b
}

func (b *breakpoint) Max(max ...any) *breakpoint {
	b.max = max
	return b
}

// breakpoint renders all the components, and uses css to hide the ones that are not needed
func (b *breakpoint) Render(w io.Writer) {
	w.Write([]byte(`<div`))
	renderAttrs(w, b.attr)
	w.Write([]byte(`>`))
	{
		w.Write([]byte(`<div class="bx-go--sm">`))
		renderAny(w, b.sm)
		w.Write([]byte(`</div>`))

		w.Write([]byte(`<div class="bx-go--md">`))
		renderAny(w, b.md)
		w.Write([]byte(`</div>`))

		w.Write([]byte(`<div class="bx-go--lg">`))
		renderAny(w, b.lg)
		w.Write([]byte(`</div>`))

		w.Write([]byte(`<div class="bx-go--xlg">`))
		renderAny(w, b.xlg)
		w.Write([]byte(`</div>`))

		w.Write([]byte(`<div class="bx-go--max">`))
		renderAny(w, b.max)
		w.Write([]byte(`</div>`))
	}
	w.Write([]byte(`</div>`))
}
