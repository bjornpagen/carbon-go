package carbon

import (
	"io"
)

type switch_ struct {
	attrs []Attr

	text     string
	selected bool
	disabled bool
}

var _ Component = (*switch_)(nil)

func Switch(text string) *switch_ {
	return &switch_{
		attrs: nil,

		text:     text,
		selected: false,
		disabled: false,
	}
}

func (s *switch_) Attr(name string, value string) *switch_ {
	s.attrs = append(s.attrs, Attr{name, value})
	return s
}

func (s *switch_) Text(text string) *switch_ {
	s.text = text
	return s
}

func (s *switch_) Disabled(disabled bool) *switch_ {
	s.disabled = disabled
	return s
}

func (s *switch_) Render(w io.Writer) {
	w.Write([]byte(`<button type="button" role="tab" tabindex="`))
	if s.selected {
		w.Write([]byte(`0" aria-selected`))
	} else {
		w.Write([]byte(`-1"`))
	}
	if s.disabled {
		w.Write([]byte(` disabled`))
	}
	w.Write([]byte(` class="bx--content-switcher-btn`))
	if s.selected {
		w.Write([]byte(` bx--content-switcher--selected`))
	}
	w.Write([]byte(`">`))
	{
		w.Write([]byte(`<span class="bx--content-switcher__label">`))
		{
			io.WriteString(w, s.text)
		}
		w.Write([]byte(`</span>`))
	}
	w.Write([]byte(`</button>`))
}
