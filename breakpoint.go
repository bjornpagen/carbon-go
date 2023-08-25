package carbon

import (
	"io"
	"strings"
)

var breakpointClass = []string{"bx-go--sm", "bx-go--md", "bx-go--lg", "bx-go--xlg", "bx-go--max"}

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
	w.Write([]byte(`<div`))
	renderAttrs(w, b.attr)
	w.Write([]byte(`>`))
	{
		// starting with "sm", we render all the components
		// breakpoints work until the next breakpoint is set, or until the end of the component slice
		// so, a single "sm" component will be rendered for all breakpoints
		// it will get the classes for every single breakpoint
		// and it will be shown for every breakpoint
		var classes []string
		lastNotNil := -1
		for i, c := range b.breakpoints {
			if c != nil {
				if lastNotNil != -1 {
					w.Write([]byte(`<div class="`))
					io.WriteString(w, strings.Join(classes, " "))
					w.Write([]byte(`">`))
					renderAny(w, b.breakpoints[lastNotNil])
					w.Write([]byte(`</div>`))
				}
				lastNotNil = i
				classes = nil
			}
			classes = append(classes, breakpointClass[i])
		}
		if lastNotNil != -1 {
			w.Write([]byte(`<div class="`))
			io.WriteString(w, strings.Join(classes, " "))
			w.Write([]byte(`">`))
			renderAny(w, b.breakpoints[lastNotNil])
			w.Write([]byte(`</div>`))
		}
	}
	w.Write([]byte(`</div>`))
}
