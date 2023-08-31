package carbon

import (
	"io"
	"slices"
)

type toggle struct {
	attrs []Attr

	size      string
	toggled   bool
	disabled  bool
	labelA    string
	labelB    string
	labelText string
	hideLabel bool
	name      string
}

var _ Component = (*toggle)(nil)

func Toggle() *toggle {
	return &toggle{
		attrs: nil,

		size:      "",
		toggled:   false,
		disabled:  false,
		labelA:    "Off",
		labelB:    "On",
		labelText: "",
		hideLabel: false,
		name:      "",
	}
}

func (t *toggle) Attr(name string, value string) *toggle {
	t.attrs = append(t.attrs, Attr{name, value})
	return t
}

func (t *toggle) Size(size string) *toggle {
	var sizes = []string{"", "sm"}
	if !slices.Contains(sizes, size) {
		panic("Size() received an invalid value")
	}
	t.size = size
	return t
}

func (t *toggle) Toggled(toggled bool) *toggle {
	t.toggled = toggled
	return t
}

func (t *toggle) Disabled(disabled bool) *toggle {
	t.disabled = disabled
	return t
}

func (t *toggle) LabelA(labelA string) *toggle {
	t.labelA = labelA
	return t
}

func (t *toggle) LabelB(labelB string) *toggle {
	t.labelB = labelB
	return t
}

func (t *toggle) LabelText(labelText string) *toggle {
	t.labelText = labelText
	return t
}

func (t *toggle) HideLabel(hideLabel bool) *toggle {
	t.hideLabel = hideLabel
	return t
}

func (t *toggle) Name(name string) *toggle {
	t.name = name
	return t
}

func (t *toggle) Render(w io.Writer) {
	id := "ccs-" + useId()
	var ariaLabel string
	if t.labelText != "" {
		ariaLabel = t.labelText
	} else if attrLabel := getAttr(t.attrs, "aria-label"); attrLabel != "" {
		ariaLabel = attrLabel
	} else {
		ariaLabel = "Toggle"
	}

	w.Write([]byte(`<div bx--form-item>`))
	{
		w.Write([]byte(`<input role="switch" type="checkbox" class="bx--toggle-input`))
		if t.size != "" {
			w.Write([]byte(` bx--toggle-input--`))
			io.WriteString(w, t.size)
		}
		w.Write([]byte(`" id="`))
		io.WriteString(w, id)
		w.Write([]byte(`"`))
		if t.toggled {
			w.Write([]byte(` checked`))
		}
		if t.disabled {
			w.Write([]byte(` disabled`))
		}
		if t.name != "" {
			w.Write([]byte(` name="`))
			io.WriteString(w, t.name)
			w.Write([]byte(`"`))
		}
		renderAttrs(w, t.attrs)
		w.Write([]byte(` />`))
		w.Write([]byte(`<label for="`))
		io.WriteString(w, id)
		w.Write([]byte(`" aria-label="`))
		io.WriteString(w, ariaLabel)
		w.Write([]byte(`" class="bx--toggle-input__label">`))
		{
			w.Write([]byte(`<span`))
			if t.hideLabel {
				w.Write([]byte(` class="bx--visually-hidden"`))
			}
			w.Write([]byte(`>`))
			{
				io.WriteString(w, t.labelText)
			}
			w.Write([]byte(`</span>`))
			w.Write([]byte(`<span class="bx--toggle__switch">`))
			{
				w.Write([]byte(`<span aria-hidden="true" class="bx--toggle__text--off">`))
				{
					io.WriteString(w, t.labelA)
				}
				w.Write([]byte(`</span>`))
				w.Write([]byte(`<span aria-hidden="true" class="bx--toggle__text--on">`))
				{
					io.WriteString(w, t.labelB)
				}
				w.Write([]byte(`</span>`))
			}
			w.Write([]byte(`</span>`))
		}
		w.Write([]byte(`</label>`))
	}
	w.Write([]byte(`</div>`))
}
