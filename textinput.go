package carbon

import (
	"fmt"
	"io"

	"golang.org/x/exp/slices"
)

type textInput struct {
	attrs []Attr

	size        string
	value       string
	placeholder string
	light       bool
	disabled    bool
	helperText  string
	id          string
	name        string
	labelText   string
	hideLabel   bool
	invalid     bool
	invalidText string
	warn        bool
	warnText    string
	required    bool
	inline      bool
	readonly    bool
	isFluid     bool
}

var _ Component = (*textInput)(nil)

func TextInput() *textInput {
	return &textInput{
		attrs: nil,

		size:        "",
		value:       "",
		placeholder: "",
		light:       false,
		disabled:    false,
		helperText:  "",
		id:          "ccs-" + useId(),
		name:        "",
		labelText:   "",
		hideLabel:   false,
		invalid:     false,
		invalidText: "",
		warn:        false,
		warnText:    "",
		required:    false,
		inline:      false,
		readonly:    false,
		isFluid:     false,
	}
}

func (t *textInput) Attr(name string, value string) Component {
	t.attrs = append(t.attrs, Attr{name, value})
	return t
}

func (t *textInput) Size(size string) *textInput {
	var sizes = []string{"", "sm", "xl"}
	if slices.Contains(sizes, size) == false {
		panic("invalid size, options are: " + fmt.Sprintf("%v", sizes))
	}
	t.size = size
	return t
}

func (t *textInput) Value(value string) *textInput {
	t.value = value
	return t
}

func (t *textInput) Placeholder(placeholder string) *textInput {
	t.placeholder = placeholder
	return t
}

func (t *textInput) Light(v bool) *textInput {
	t.light = v
	return t
}

func (t *textInput) Disabled(v bool) *textInput {
	t.disabled = v
	return t
}

func (t *textInput) HelperText(helperText string) *textInput {
	t.helperText = helperText
	return t
}

func (t *textInput) Name(name string) *textInput {
	t.name = name
	return t
}

func (t *textInput) LabelText(labelText string) *textInput {
	t.labelText = labelText
	return t
}

func (t *textInput) HideLabel(v bool) *textInput {
	t.hideLabel = v
	return t
}

func (t *textInput) Invalid(v bool) *textInput {
	t.invalid = v
	return t
}

func (t *textInput) InvalidText(invalidText string) *textInput {
	t.invalidText = invalidText
	return t
}

func (t *textInput) Warn(v bool) *textInput {
	t.warn = v
	return t
}

func (t *textInput) WarnText(warnText string) *textInput {
	t.warnText = warnText
	return t
}

func (t *textInput) Required(v bool) *textInput {
	t.required = v
	return t
}

func (t *textInput) Inline(v bool) *textInput {
	t.inline = v
	return t
}

func (t *textInput) Readonly(v bool) *textInput {
	t.readonly = v
	return t
}

func (t *textInput) Fluid(v bool) *textInput {
	t.isFluid = v
	return t
}

func (t *textInput) Render(w io.Writer) {
	error := t.invalid && !t.readonly
	helperId := "helper-" + t.id
	errorId := "error-" + t.id
	warnId := "warn-" + t.id

	renderLabelText := renderLabelTextClosure(t.id, t.hideLabel, t.disabled, t.inline, t.size, t.labelText)
	renderHelperText := renderHelperTextClosure(t.disabled, t.inline, t.helperText)

	w.Write([]byte(`<div class="bx--form-item bx--text-input-wrapper`))
	if t.inline {
		w.Write([]byte(` bx--text-input-wrapper--inline`))
	}
	if t.light {
		w.Write([]byte(` bx--text-input-wrapper--light`))
	}
	if t.readonly {
		w.Write([]byte(` bx--text-input-wrapper--readonly`))
	}
	w.Write([]byte(`">`))

	if t.inline {
		w.Write([]byte(`<div class="bx--text-input__label-helper-wrapper">`))
		if t.labelText != "" {
			renderLabelText(w)
		}
		if !t.isFluid && t.helperText != "" {
			renderHelperText(w)
		}
		w.Write([]byte(`</div>`))
	}

	if !t.inline && t.labelText != "" {
		renderLabelText(w)
	}

	w.Write([]byte(`<div class="bx--text-input__field-outer-wrapper`))
	if t.inline {
		w.Write([]byte(` bx--text-input__field-outer-wrapper--inline`))
	}
	w.Write([]byte(`">`))

	{
		w.Write([]byte(`<div`))
		if error {
			w.Write([]byte(` data-invalid`))
		}
		if t.warn {
			w.Write([]byte(` data-warn`))
		}
		w.Write([]byte(` class="bx--text-input__field-wrapper`))
		if !t.invalid && t.warn {
			w.Write([]byte(` bx--text-input__field-wrapper--warning`))
		}
		w.Write([]byte(`">`))

		{
			if t.readonly {
				EditOff().Attr("class", "bx--text-input__readonly-icon").Render(w)
			} else {
				if t.invalid {
					WarningFilled().Attr("class", "bx--text-input__invalid-icon").Render(w)
				}
				if !t.invalid && t.warn {
					WarningAltFilled().Attr("class", "bx--text-input__invalid-icon bx--text-input__invalid-icon--warning").Render(w)
				}
			}
			w.Write([]byte(`<input`))
			if error {
				w.Write([]byte(` data-invalid aria-invalid`))
			}
			if t.warn {
				w.Write([]byte(` data-warn`))
			}
			describedBy := ternary(error, errorId, ternary(t.warn, warnId, ternary(t.helperText != "", helperId, "")))
			if describedBy != "" {
				w.Write([]byte(` aria-describedby="`))
				w.Write(yoloBytesUnsafe(describedBy))
				w.Write([]byte(`"`))
			}
			if t.disabled {
				w.Write([]byte(` disabled`))
			}
			w.Write([]byte(` id="`))
			w.Write(yoloBytesUnsafe(t.id))
			w.Write([]byte(`"`))
			if t.name != "" {
				w.Write([]byte(` name="`))
				w.Write(yoloBytesUnsafe(t.name))
				w.Write([]byte(`"`))
			}
			if t.placeholder != "" {
				w.Write([]byte(` placeholder="`))
				w.Write(yoloBytesUnsafe(t.placeholder))
				w.Write([]byte(`"`))
			}
			if t.value != "" {
				w.Write([]byte(` value="`))
				w.Write(yoloBytesUnsafe(t.value))
				w.Write([]byte(`"`))
			}
			if t.required {
				w.Write([]byte(` required`))
			}
			if t.readonly {
				w.Write([]byte(` readonly`))
			}
			w.Write([]byte(` class="bx--text-input`))
			if t.light {
				w.Write([]byte(` bx--text-input--light`))
			}
			if error {
				w.Write([]byte(` bx--text-input--invalid`))
			}
			if t.warn {
				w.Write([]byte(` bx--text-input--warning`))
			}
			if t.size == "sm" || t.size == "xl" {
				w.Write([]byte(` bx--text-input--`))
				w.Write(yoloBytesUnsafe(t.size))
			}
			w.Write([]byte(`"`))
			renderAttrs(w, t.attrs)
			w.Write([]byte(`/>`))
			if t.isFluid {
				w.Write([]byte(`<hr class="bx--text-input__divider" />`))
			}
			if t.isFluid && !t.inline && t.invalid {
				w.Write([]byte(`<div class="bx--form-requirement" id="`))
				w.Write(yoloBytesUnsafe(errorId))
				w.Write([]byte(`">`))
				w.Write(yoloBytesUnsafe(t.invalidText))
				w.Write([]byte(`</div>`))
			}
			if t.isFluid && !t.inline && t.warn {
				w.Write([]byte(`<div class="bx--form-requirement" id="`))
				w.Write(yoloBytesUnsafe(warnId))
				w.Write([]byte(`">`))
				w.Write(yoloBytesUnsafe(t.warnText))
				w.Write([]byte(`</div>`))
			}
		}
		w.Write([]byte(`</div>`))

		if !t.invalid && !t.warn && !t.isFluid && !t.inline && t.helperText != "" {
			w.Write([]byte(`<div id="`))
			w.Write(yoloBytesUnsafe(helperId))
			w.Write([]byte(`" class="bx--form__helper-text`))
			if t.disabled {
				w.Write([]byte(` bx--form__helper-text--disabled`))
			}
			if t.inline {
				w.Write([]byte(` bx--form__helper-text--inline`))
			}
			w.Write([]byte(`">`))
			w.Write(yoloBytesUnsafe(t.helperText))
			w.Write([]byte(`</div>`))
		}
		if !t.isFluid && t.invalid {
			w.Write([]byte(`<div class="bx--form-requirement" id="`))
			w.Write(yoloBytesUnsafe(errorId))
			w.Write([]byte(`">`))
			w.Write(yoloBytesUnsafe(t.invalidText))
			w.Write([]byte(`</div>`))
		}
		if !t.isFluid && !t.invalid && t.warn {
			w.Write([]byte(`<div class="bx--form-requirement" id="`))
			w.Write(yoloBytesUnsafe(warnId))
			w.Write([]byte(`">`))
			w.Write(yoloBytesUnsafe(t.warnText))
			w.Write([]byte(`</div>`))
		}
	}
	w.Write([]byte(`</div>`))
	w.Write([]byte(`</div>`))
}

func renderLabelTextClosure(forId string, hideLabel, disabled, inline bool, size, labelText string) func(io.Writer) {
	return func(w io.Writer) {
		w.Write([]byte(`<label for="`))
		w.Write(yoloBytesUnsafe(forId))
		w.Write([]byte(`" class="bx--label`))
		if hideLabel {
			w.Write([]byte(` bx--visually-hidden`))
		}
		if disabled {
			w.Write([]byte(` bx--label--disabled`))
		}
		if inline {
			w.Write([]byte(` bx--label--inline`))
		}
		if size == "sm" || size == "xl" {
			w.Write([]byte(` bx--label--inline--`))
			w.Write(yoloBytesUnsafe(size))
		}
		w.Write([]byte(`">`))
		w.Write(yoloBytesUnsafe(labelText))
		w.Write([]byte(`</label>`))
	}
}

func renderHelperTextClosure(disabled, inline bool, helperText string) func(io.Writer) {
	return func(w io.Writer) {
		w.Write([]byte(`<div class="bx--form__helper-text`))
		if disabled {
			w.Write([]byte(` bx--form__helper-text--disabled`))
		}
		if inline {
			w.Write([]byte(` bx--form__helper-text--inline`))
		}
		w.Write([]byte(`">`))
		w.Write(yoloBytesUnsafe(helperText))
		w.Write([]byte(`</div>`))
	}
}
