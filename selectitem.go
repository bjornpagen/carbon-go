package carbon

import "io"

type selectItem struct {
	attrs []Attr

	value    string
	text     string
	hidden   bool
	disabled bool
	selected bool
}

var _ Component = (*selectItem)(nil)

func SelectItem() *selectItem {
	return &selectItem{
		attrs: nil,

		value:    "",
		text:     "",
		hidden:   false,
		disabled: false,
		selected: false,
	}
}

func (s *selectItem) Attr(name string, value string) Component {
	s.attrs = append(s.attrs, Attr{name, value})
	return s
}

func (s *selectItem) Value(value string) *selectItem {
	s.value = value
	return s
}

func (s *selectItem) Text(text string) *selectItem {
	s.text = text
	return s
}

func (s *selectItem) Hidden(hidden bool) *selectItem {
	s.hidden = hidden
	return s
}

func (s *selectItem) Disabled(disabled bool) *selectItem {
	s.disabled = disabled
	return s
}

func (s *selectItem) Render(w io.Writer) {
	w.Write([]byte(`<option`))
	if s.value != "" {
		w.Write([]byte(` value="`))
		w.Write(yoloBytesUnsafe(s.value))
		w.Write([]byte(`"`))
	}
	if s.disabled {
		w.Write([]byte(` disabled`))
	}
	if s.hidden {
		w.Write([]byte(` hidden`))
	}
	if s.selected {
		w.Write([]byte(` selected`))
	}
	w.Write([]byte(` class="bx--select-option">`))
	{
		if s.text != "" {
			w.Write(yoloBytesUnsafe(s.text))
		} else {
			w.Write(yoloBytesUnsafe(s.value))
		}
	}
	w.Write([]byte(`</option>`))
}
