package carbon

import (
	"io"
	"slices"
	"strings"
)

type passwordInput struct {
	attrs []Attr

	size              string
	value             string
	type_             string
	placeholder       string
	hidePasswordLabel string
	showPasswordLabel string
	tooltipAlignment  string
	tooltipPosition   string
	light             bool
	disabled          bool
	helperText        string
	labelText         string
	hideLabel         bool
	invalid           bool
	invalidText       string
	warn              bool
	warnText          string
	inline            bool
	name              string
	isFluid           bool
}

var _ Component = (*passwordInput)(nil)

func PasswordInput(labelText string) *passwordInput {
	return &passwordInput{
		attrs: nil,

		size:              "",
		value:             "",
		type_:             "password",
		placeholder:       "",
		hidePasswordLabel: "Hide password",
		showPasswordLabel: "Show password",
		tooltipAlignment:  "center",
		tooltipPosition:   "bottom",
		light:             false,
		disabled:          false,
		helperText:        "",
		labelText:         labelText,
		hideLabel:         false,
		invalid:           false,
		invalidText:       "",
		warn:              false,
		warnText:          "",
		inline:            false,
		name:              "",
		isFluid:           false,
	}
}

func (p *passwordInput) Attr(name string, value string) *passwordInput {
	p.attrs = append(p.attrs, Attr{name, value})
	return p
}

func (p *passwordInput) Size(size string) *passwordInput {
	sizes := []string{"", "sm", "xl"}
	if !slices.Contains(sizes, size) {
		panic("size must be one of " + strings.Join(sizes, ", "))
	}
	p.size = size
	return p
}

func (p *passwordInput) Value(value string) *passwordInput {
	p.value = value
	return p
}

func (p *passwordInput) Type(type_ string) *passwordInput {
	types := []string{"password", "text"}
	if !slices.Contains(types, type_) {
		panic("type must be one of " + strings.Join(types, ", "))
	}
	p.type_ = type_
	return p
}

func (p *passwordInput) Placeholder(placeholder string) *passwordInput {
	p.placeholder = placeholder
	return p
}

func (p *passwordInput) HidePasswordLabel(hidePasswordLabel string) *passwordInput {
	p.hidePasswordLabel = hidePasswordLabel
	return p
}

func (p *passwordInput) ShowPasswordLabel(showPasswordLabel string) *passwordInput {
	p.showPasswordLabel = showPasswordLabel
	return p
}

func (p *passwordInput) TooltipAlignment(tooltipAlignment string) *passwordInput {
	tooltipAlignments := []string{"start", "center", "end"}
	if !slices.Contains(tooltipAlignments, tooltipAlignment) {
		panic("tooltipAlignment must be one of " + strings.Join(tooltipAlignments, ", "))
	}
	p.tooltipAlignment = tooltipAlignment
	return p
}

func (p *passwordInput) TooltipPosition(tooltipPosition string) *passwordInput {
	tooltipPositions := []string{"top", "right", "bottom", "left"}
	if !slices.Contains(tooltipPositions, tooltipPosition) {
		panic("tooltipPosition must be one of " + strings.Join(tooltipPositions, ", "))
	}
	p.tooltipPosition = tooltipPosition
	return p
}

func (p *passwordInput) Light(v bool) *passwordInput {
	p.light = v
	return p
}

func (p *passwordInput) Disabled(v bool) *passwordInput {
	p.disabled = v
	return p
}

func (p *passwordInput) HelperText(helperText string) *passwordInput {
	p.helperText = helperText
	return p
}

func (p *passwordInput) HideLabel(v bool) *passwordInput {
	p.hideLabel = v
	return p
}

func (p *passwordInput) Invalid(v bool) *passwordInput {
	p.invalid = v
	return p
}

func (p *passwordInput) InvalidText(invalidText string) *passwordInput {
	p.invalidText = invalidText
	return p
}

func (p *passwordInput) Warn(v bool) *passwordInput {
	p.warn = v
	return p
}

func (p *passwordInput) WarnText(warnText string) *passwordInput {
	p.warnText = warnText
	return p
}

func (p *passwordInput) Inline(v bool) *passwordInput {
	p.inline = v
	return p
}

func (p *passwordInput) Name(name string) *passwordInput {
	p.name = name
	return p
}

func (p *passwordInput) IsFluid(v bool) *passwordInput {
	p.isFluid = v
	return p
}

func (p *passwordInput) Render(w io.Writer) {
	id := "ccs-" + useId()
	helperId := "helper-" + id
	errorId := "error-" + id
	warnId := "warn-" + id

	renderLabelText := renderLabelTextClosure(id, p.hideLabel, p.disabled, p.inline, p.size, p.labelText)
	renderHelperText := renderHelperTextClosure(p.disabled, p.inline, p.helperText)

	w.Write([]byte(`<div class="bx--form-item bx--text-input-wrapper`))
	if !p.isFluid {
		w.Write([]byte(` bx--password-input-wrapper`))
	}
	if p.light {
		w.Write([]byte(` bx--text-input-wrapper--light`))
	}
	if p.inline {
		w.Write([]byte(` bx--text-input-wrapper--inline`))
	}
	w.Write([]byte(`">`))
	{
		if p.inline {
			renderLabelText(w)
			if !p.isFluid && p.helperText != "" {
				renderHelperText(w)
			}
		}
		if !p.inline && p.labelText != "" {
			renderLabelText(w)
		}
		w.Write([]byte(`<div class="bx--text-input__field-outer-wrapper`))
		if p.inline {
			w.Write([]byte(` bx--text-input__field-outer-wrapper--inline`))
		}
		w.Write([]byte(`">`))
		{
			w.Write([]byte(`<div class="bx--text-input__field-wrapper`))
			if p.warn {
				w.Write([]byte(` bx--text-input__field-wrapper--warning`))
			}
			w.Write([]byte(`"`))
			if p.invalid {
				w.Write([]byte(` data-invalid`))
			}
			w.Write([]byte(`>`))
			{
				if p.invalid {
					WarningFilled().Attr("class", "bx--text-input__invalid-icon").Render(w)
				}
				if !p.invalid && p.warn {
					WarningAltFilled().Attr("class", "bx--text-input__invalid-icon bx--text-input__invalid-icon--warning").Render(w)
				}
				w.Write([]byte(`<input`))
				if p.invalid {
					w.Write([]byte(` data-invalid aria-invalid`))
				}
				describedBy := ternary(p.invalid, errorId, ternary(p.warn, warnId, ternary(p.helperText != "", helperId, "")))
				if describedBy != "" {
					w.Write([]byte(` aria-describedby="`))
					io.WriteString(w, describedBy)
					w.Write([]byte(`"`))
				}
				w.Write([]byte(` id="`))
				io.WriteString(w, id)
				w.Write([]byte(`"`))
				if p.name != "" {
					w.Write([]byte(` name="`))
					io.WriteString(w, p.name)
					w.Write([]byte(`"`))
				}
				if p.placeholder != "" {
					w.Write([]byte(` placeholder="`))
					io.WriteString(w, p.placeholder)
					w.Write([]byte(`"`))
				}
				w.Write([]byte(` type="`))
				io.WriteString(w, p.type_)
				w.Write([]byte(`"`))
				if p.value != "" {
					w.Write([]byte(` value="`))
					io.WriteString(w, p.value)
					w.Write([]byte(`"`))
				}
				if p.disabled {
					w.Write([]byte(` disabled`))
				}
				w.Write([]byte(` class="bx--text-input bx--password-input`))
				if p.light {
					w.Write([]byte(` bx--text-input--light`))
				}
				if p.invalid {
					w.Write([]byte(` bx--text-input--invalid`))
				}
				if p.warn {
					w.Write([]byte(` bx--text-input--warning`))
				}
				if p.size == "sm" || p.size == "xl" {
					w.Write([]byte(` bx--text-input--`))
					io.WriteString(w, p.size)
				}
				w.Write([]byte(`"`))
				renderAttrs(w, p.attrs)
				w.Write([]byte(` />`))

				if p.isFluid && p.invalid {
					w.Write([]byte(`<hr class="bx--text-input__divider"/><div class="bx--form-requirement" id="`))
					io.WriteString(w, errorId)
					w.Write([]byte(`">`))
					io.WriteString(w, p.invalidText)
					w.Write([]byte(`</div>`))
				} else {
					w.Write([]byte(`<button type="button" disabled class="bx--text-input--password__visibility__toggle bx--btn bx--btn--icon-only bx--tooltip__trigger bx--tooltip--a11y`))
					if p.disabled {
						w.Write([]byte(` disabled`))
					}
					w.Write([]byte(` bx--tooltip--`))
					io.WriteString(w, p.tooltipPosition)
					w.Write([]byte(` bx--tooltip--align-`))
					io.WriteString(w, p.tooltipAlignment)
					w.Write([]byte(`">`))
					{
						if !p.disabled {
							w.Write([]byte(`<span class="bx--assistive-text">`))
							if p.type_ == "text" {
								io.WriteString(w, p.hidePasswordLabel)
							} else {
								io.WriteString(w, p.showPasswordLabel)
							}
							w.Write([]byte(`</span>`))
						}
						if p.type_ == "text" {
							ViewOff().Attr("class", "bx--icon-visibility-off").Render(w)
						} else {
							View().Attr("class", "bx--icon-visibility-on").Render(w)
						}
					}
					w.Write([]byte(`</button>`))
				}
			}
			w.Write([]byte(`</div>`))
			if !p.isFluid && p.invalid {
				w.Write([]byte(`<div class="bx--form-requirement" id="`))
				io.WriteString(w, errorId)
				w.Write([]byte(`">`))
				io.WriteString(w, p.invalidText)
				w.Write([]byte(`</div>`))
			}
			if !p.invalid && !p.warn && !p.isFluid && !p.inline && p.helperText != "" {
				w.Write([]byte(`<div class="bx--form__helper-text`))
				if p.disabled {
					w.Write([]byte(` bx--form__helper-text--disabled`))
				}
				if p.inline {
					w.Write([]byte(` bx--form__helper-text--inline`))
				}
				w.Write([]byte(`">`))
				io.WriteString(w, p.helperText)
				w.Write([]byte(`</div>`))
			}
			if !p.isFluid && !p.invalid && p.warn {
				w.Write([]byte(`<div class="bx--form-requirement" id="`))
				io.WriteString(w, warnId)
				w.Write([]byte(`">`))
				io.WriteString(w, p.warnText)
				w.Write([]byte(`</div>`))
			}
		}
		w.Write([]byte(`</div>`))
	}
	w.Write([]byte(`</div>`))
}
