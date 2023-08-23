package carbon

import (
	"io"
	"slices"
)

type select_ struct {
	children []*selectItem
	attrs    []Attr

	selected    string
	size        string
	inline      bool
	light       bool
	disabled    bool
	id          string
	name        string
	invalid     bool
	invalidText string
	warn        bool
	warnText    string
	helperText  string
	noLabel     bool
	labelText   string
	hideLabel   bool
	required    bool
}

var _ Component = (*select_)(nil)

func Select(s ...*selectItem) *select_ {
	return &select_{
		children: s,
		attrs:    nil,

		selected:    "",
		size:        "",
		inline:      false,
		light:       false,
		disabled:    false,
		id:          "ccs-" + useId(),
		name:        "",
		invalid:     false,
		invalidText: "",
		warn:        false,
		warnText:    "",
		helperText:  "",
		noLabel:     false,
		labelText:   "",
		hideLabel:   false,
		required:    false,
	}
}

func (s *select_) Attr(name string, value string) Component {
	s.attrs = append(s.attrs, Attr{name, value})
	return s
}

func (s *select_) Selected(selected string) *select_ {
	s.selected = selected
	return s
}

func (s *select_) Size(size string) *select_ {
	var sizes = []string{"", "sm", "xl"}
	if !slices.Contains(sizes, size) {
		panic("invalid size")
	}
	s.size = size
	return s
}

func (s *select_) Inline(inline bool) *select_ {
	s.inline = inline
	return s
}

func (s *select_) Light(light bool) *select_ {
	s.light = light
	return s
}

func (s *select_) Disabled(disabled bool) *select_ {
	s.disabled = disabled
	return s
}

func (s *select_) Id(id string) *select_ {
	s.id = id
	return s
}

func (s *select_) Name(name string) *select_ {
	s.name = name
	return s
}

func (s *select_) Invalid(invalid bool) *select_ {
	s.invalid = invalid
	return s
}

func (s *select_) InvalidText(invalidText string) *select_ {
	s.invalidText = invalidText
	return s
}

func (s *select_) Warn(warn bool) *select_ {
	s.warn = warn
	return s
}

func (s *select_) WarnText(warnText string) *select_ {
	s.warnText = warnText
	return s
}

func (s *select_) HelperText(helperText string) *select_ {
	s.helperText = helperText
	return s
}

func (s *select_) NoLabel(noLabel bool) *select_ {
	s.noLabel = noLabel
	return s
}

func (s *select_) LabelText(labelText string) *select_ {
	s.labelText = labelText
	return s
}

func (s *select_) HideLabel(hideLabel bool) *select_ {
	s.hideLabel = hideLabel
	return s
}

func (s *select_) Required(required bool) *select_ {
	s.required = required
	return s
}

func (s *select_) Render(w io.Writer) {
	errorId := `error-` + s.id
	renderSelect := renderSelectClosure(s.invalid, s.disabled, s.required, errorId, s.id, s.name, s.size, s.selected, s.attrs, s.children)

	w.Write([]byte(`<div class="bx--form-item"><div class="bx--select`))
	if s.inline {
		w.Write([]byte(` bx--select--inline`))
	}
	if s.light {
		w.Write([]byte(` bx--select--light`))
	}
	if s.invalid {
		w.Write([]byte(` bx--select--invalid`))
	}
	if s.disabled {
		w.Write([]byte(` bx--select--disabled`))
	}
	if s.warn {
		w.Write([]byte(` bx--select--warning`))
	}
	w.Write([]byte(`">`))
	{
		if !s.noLabel {
			w.Write([]byte(`<label for="`))
			w.Write(yoloBytesUnsafe(s.id))
			w.Write([]byte(`" class="bx--label`))
			if s.hideLabel {
				w.Write([]byte(` bx--visually-hidden`))
			}
			if s.disabled {
				w.Write([]byte(` bx--label--disabled`))
			}
			w.Write([]byte(`">`))
			{
				w.Write(yoloBytesUnsafe(s.labelText))
			}
			w.Write([]byte(`</label>`))
		}
		if s.inline {
			w.Write([]byte(`<div class="bx--select-input--inline__wrapper">`))
			{
				w.Write([]byte(`<div class="bx--select-input__wrapper"`))
				if s.invalid {
					w.Write([]byte(` data-invalid`))
				}
				w.Write([]byte(`>`))
				{
					renderSelect(w)

					ChevronDown().Attr("class", "bx--select__arrow").Render(w)

					if s.invalid {
						WarningFilled().Attr("class", "bx--select__invalid-icon").Render(w)
					}
				}
				w.Write([]byte(`</div>`))
				if s.invalid {
					w.Write([]byte(`<div class="bx--form-requirement" id="`))
					w.Write(yoloBytesUnsafe(errorId))
					w.Write([]byte(`">`))
					{
						w.Write(yoloBytesUnsafe(s.invalidText))
					}
					w.Write([]byte(`</div>`))
				}
				w.Write([]byte(`</div>`))
				if !s.invalid && !s.warn && s.helperText != "" {
					w.Write([]byte(`<div class="bx--form__helper-text`))
					if s.disabled {
						w.Write([]byte(` bx--form__helper-text--disabled`))
					}
					w.Write([]byte(`">`))
					{
						w.Write(yoloBytesUnsafe(s.helperText))
					}
					w.Write([]byte(`</div>`))
				}
			}
			w.Write([]byte(`</div>`))
		}
		if !s.inline {
			w.Write([]byte(`<div class="bx--select-input__wrapper"`))
			if s.invalid {
				w.Write([]byte(` data-invalid`))
			}
			w.Write([]byte(`>`))
			{
				renderSelect(w)

				ChevronDown().Attr("class", "bx--select__arrow").Render(w)

				if s.invalid {
					WarningFilled().Attr("class", "bx--select__invalid-icon").Render(w)
				}
				if !s.invalid && s.warn {
					WarningAltFilled().Attr("class", "bx--select__invalid-icon bx--select__invalid-icon--warning").Render(w)
				}
			}
			w.Write([]byte(`</div>`))
			if !s.invalid && s.helperText != "" {
				w.Write([]byte(`<div class="bx--form__helper-text`))
				if s.disabled {
					w.Write([]byte(` bx--form__helper-text--disabled`))
				}
				w.Write([]byte(`">`))
				{
					w.Write(yoloBytesUnsafe(s.helperText))
				}
				w.Write([]byte(`</div>`))
			}
			if s.invalid {
				w.Write([]byte(`<div id="`))
				w.Write(yoloBytesUnsafe(errorId))
				w.Write([]byte(`" class="bx--form-requirement">`))
				{
					w.Write(yoloBytesUnsafe(s.invalidText))
				}
				w.Write([]byte(`</div>`))
			}
			if !s.invalid && s.warn {
				w.Write([]byte(`<div id="`))
				w.Write(yoloBytesUnsafe(errorId))
				w.Write([]byte(`" class="bx--form-requirement">`))
				{
					w.Write(yoloBytesUnsafe(s.warnText))
				}
				w.Write([]byte(`</div>`))
			}
		}
	}
	w.Write([]byte(`</div></div>`))
}

func renderSelectClosure(invalid, disabled, required bool, errorId, id, name, size, selected string, attrs []Attr, children []*selectItem) func(io.Writer) {
	return func(w io.Writer) {
		w.Write([]byte(`<select`))
		if invalid {
			w.Write([]byte(` aria-describedby="`))
			w.Write(yoloBytesUnsafe(errorId))
			w.Write([]byte(`" aria-invalid`))
		}
		if disabled {
			w.Write([]byte(` disabled`))
		}
		if required {
			w.Write([]byte(` required`))
		}
		w.Write([]byte(` id="`))
		w.Write(yoloBytesUnsafe(id))
		w.Write([]byte(`" name="`))
		w.Write(yoloBytesUnsafe(name))
		w.Write([]byte(`" class="bx--select-input`))
		switch size {
		case "sm":
			w.Write([]byte(` bx--select-input--sm`))
		case "xl":
			w.Write([]byte(` bx--select-input--xl`))
		}
		renderAttrs(w, attrs)
		w.Write([]byte(`">`))
		{
			for _, child := range children {
				if child.value == selected {
					child.selected = true
				} else {
					child.selected = false
				}
				child.Render(w)
			}
		}
		w.Write([]byte(`</select>`))
	}
}
