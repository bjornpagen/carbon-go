package carbon

import "io"

var breakpointClass = []string{"sm", "md", "lg", "xlg", "max"}

type breakpoint struct {
	attr []Attr

	breakpoints [5]any
}

func Breakpoint(sm ...any) *breakpoint {
	return &breakpoint{
		attr: nil,

		breakpoints: [5]any{sm, nil, nil, nil, nil},
	}
}

func (b *breakpoint) Attr(name string, value string) Component {
	b.attr = append(b.attr, Attr{name, value})
	return b
}

func (b *breakpoint) Md(md ...any) *breakpoint {
	b.breakpoints[1] = md
	return b
}

func (b *breakpoint) Lg(lg ...any) *breakpoint {
	b.breakpoints[2] = lg
	return b
}

func (b *breakpoint) Xlg(xlg ...any) *breakpoint {
	b.breakpoints[3] = xlg
	return b
}

func (b *breakpoint) Max(max ...any) *breakpoint {
	b.breakpoints[4] = max
	return b
}

// breakpoint renders all the components, and uses css to hide the ones that are not needed
func (b *breakpoint) Render(w io.Writer) {
	i := 0
	w.Write([]byte(`<div`))
	renderAttrs(w, b.attr)
	w.Write([]byte(`>`))
	{
		for j := range b.breakpoints {
			if b.breakpoints[j] != nil {
				i = j
			}
			w.Write([]byte(`<div class="bx-go--`))
			w.Write([]byte(breakpointClass[j]))
			w.Write([]byte(`">`))
			{
				renderAny(w, b.breakpoints[i])
			}
			w.Write([]byte(`</div>`))
		}
	}
	w.Write([]byte(`</div>`))
}
