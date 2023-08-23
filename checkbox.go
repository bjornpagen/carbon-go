package carbon

import "io"

type checkbox struct {
	attrs         []Attr
	value         string
	checked       bool
	indeterminate bool
	required      bool
	readonly      bool
	disabled      bool
	labelText     string
	hideLabel     bool
	name          string
	title         string
	id            string
}

var _ Component = (*checkbox)(nil)

func Checkbox() *checkbox {
	return &checkbox{
		attrs:         nil,
		value:         "",
		checked:       false,
		indeterminate: false,
		required:      false,
		readonly:      false,
		disabled:      false,
		labelText:     "",
		hideLabel:     false,
		name:          "",
		title:         "",
		id:            "ccs-" + useId(),
	}
}

func (c *checkbox) Attr(name string, value string) Component {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *checkbox) Value(value string) *checkbox {
	c.value = value
	return c
}

func (c *checkbox) Checked(checked bool) *checkbox {
	c.checked = checked
	return c
}

func (c *checkbox) Indeterminate(indeterminate bool) *checkbox {
	panic("not implemented")
	// c.indeterminate = indeterminate
	// return c
}

func (c *checkbox) Required(required bool) *checkbox {
	c.required = required
	return c
}

func (c *checkbox) Readonly(readonly bool) *checkbox {
	c.readonly = readonly
	return c
}

func (c *checkbox) Disabled(disabled bool) *checkbox {
	c.disabled = disabled
	return c
}

func (c *checkbox) LabelText(labelText string) *checkbox {
	c.labelText = labelText
	return c
}

func (c *checkbox) HideLabel(hideLabel bool) *checkbox {
	c.hideLabel = hideLabel
	return c
}

func (c *checkbox) Name(name string) *checkbox {
	c.name = name
	return c
}

func (c *checkbox) Title(title string) *checkbox {
	c.title = title
	return c
}

func (c *checkbox) Id(id string) *checkbox {
	c.id = id
	return c
}

func (c *checkbox) Render(w io.Writer) {
	w.Write([]byte(`<div class="bx--form-item bx--checkbox-wrapper">`))
	{
		w.Write([]byte(`<input type="checkbox" value="`))
		w.Write(yoloBytesUnsafe(c.value))
		w.Write([]byte(`"`))
		if c.checked {
			w.Write([]byte(` checked`))
		}
		if c.disabled {
			w.Write([]byte(` disabled`))
		}
		w.Write([]byte(` id="`))
		w.Write(yoloBytesUnsafe(c.id))
		w.Write([]byte(`"`))
		if c.name != "" {
			w.Write([]byte(` name="`))
			w.Write(yoloBytesUnsafe(c.name))
			w.Write([]byte(`"`))
		}
		if c.required {
			w.Write([]byte(` required`))
		}
		if c.readonly {
			w.Write([]byte(` readonly`))
		}
		w.Write([]byte(` class="bx--checkbox"`))
		renderAttrs(w, c.attrs)
		w.Write([]byte(`/><label for="`))
		w.Write(yoloBytesUnsafe(c.id))
		w.Write([]byte(`"`))
		if c.title != "" {
			w.Write([]byte(` title="`))
			w.Write(yoloBytesUnsafe(c.title))
			w.Write([]byte(`"`))
		}
		w.Write([]byte(` class="bx--checkbox-label">`))
		{
			w.Write([]byte(`<span class="bx--checkbox-label-text`))
			if c.hideLabel {
				w.Write([]byte(` bx--visually-hidden`))
			}
			w.Write([]byte(`">`))
			{
				w.Write(yoloBytesUnsafe(c.labelText))
			}
			w.Write([]byte(`</span>`))
		}
		w.Write([]byte(`</label>`))
	}
	w.Write([]byte(`</div>`))
}
