package carbon

import (
	"io"

	"github.com/valyala/bytebufferpool"
	"golang.org/x/exp/slices"
)

type button struct {
	children Component
	attrs    []Attr

	kind             string
	size             string
	expressive       bool
	isSelected       bool
	icon             Component
	iconDescription  string
	tooltipAlignment string
	tooltipPosition  string
	skeleton         bool
	disabled         bool
	href             string
	type_            string
}

var _ Component = (*button)(nil)

func Button(children Component) *button {
	return &button{
		children: children,
		attrs:    nil,

		kind:             "primary",
		size:             "default",
		expressive:       false,
		isSelected:       false,
		icon:             nil,
		iconDescription:  "",
		tooltipAlignment: "",
		tooltipPosition:  "",
		skeleton:         false,
		disabled:         false,
		href:             "",
		type_:            "button",
	}
}

func (b *button) Attrs(attrs ...Attr) *button {
	b.attrs = attrs
	return b
}

func (b *button) Kind(kind string) *button {
	var kinds = []string{"primary", "secondary", "tertiary", "ghost", "danger", "danger-tertiary", "danger-ghost"}
	if slices.Contains(kinds, kind) == false {
		panic("invalid kind")
	}

	b.kind = kind
	return b
}

func (b *button) Size(size string) *button {
	var sizes = []string{"default", "field", "sm", "lg", "xl"}
	if slices.Contains(sizes, size) == false {
		panic("invalid size")
	}

	b.size = size
	return b
}

func (b *button) Expressive(v bool) *button {
	b.expressive = v
	return b
}

func (b *button) Selected(v bool) *button {
	b.isSelected = v
	return b
}

func (b *button) Icon(icon Component) *button {
	b.icon = icon
	return b
}

func (b *button) IconDescription(iconDescription string) *button {
	b.iconDescription = iconDescription
	return b
}

func (b *button) TooltipAlignment(tooltipAlignment string) *button {
	b.tooltipAlignment = tooltipAlignment
	return b
}

func (b *button) TooltipPosition(tooltipPosition string) *button {
	b.tooltipPosition = tooltipPosition
	return b
}

func (b *button) Skeleton(v bool) *button {
	b.skeleton = v
	return b
}

func (b *button) Disabled(v bool) *button {
	b.disabled = v
	return b
}

func (b *button) Href(href string) *button {
	b.href = href
	return b
}

func (b *button) Type(type_ string) *button {
	b.type_ = type_
	return b
}

func (b *button) Render(w io.Writer) {
	hasIconOnly := b.icon != nil && b.children == nil
	ariaPressed := ternary(hasIconOnly && b.kind == "ghost" && b.href == "", b.isSelected, false)

	class := bytebufferpool.Get()
	defer bytebufferpool.Put(class)

	class.WriteString(`bx--btn`)
	if b.expressive {
		class.WriteString(` bx--btn--expressive`)
	}
	if b.size == "sm" && !b.expressive {
		class.WriteString(` bx--btn--sm`)
	}
	if !(b.size == "field" && !b.expressive) && b.size == "md" && !b.expressive {
		class.WriteString(` bx--btn--md`)
	}
	if b.size == "field" {
		class.WriteString(` bx--btn--field`)
	}
	if b.size == "sm" {
		class.WriteString(` bx--btn--sm`)
	}
	if b.size == "lg" {
		class.WriteString(` bx--btn--lg`)
	}
	if b.size == "xl" {
		class.WriteString(` bx--btn--xl`)
	}
	class.WriteString(` bx--btn--`)
	class.WriteString(b.kind)
	if b.disabled {
		class.WriteString(` bx--btn--disabled`)
	}
	if hasIconOnly {
		class.WriteString(` bx--btn--icon-only`)
		class.WriteString(` bx--tooltip__trigger`)
		class.WriteString(` bx--tooltip--a11y`)
	}
	if hasIconOnly && b.tooltipPosition != "" {
		class.WriteString(` bx--btn--icon-only--`)
		class.WriteString(b.tooltipPosition)
	}
	if hasIconOnly && b.tooltipAlignment != "" {
		class.WriteString(` bx--tooltip--align-`)
		class.WriteString(b.tooltipAlignment)
	}
	if hasIconOnly && b.isSelected && b.kind == "ghost" {
		class.WriteString(` bx--btn--selected`)
	}

	attrs := slices.Clone(b.attrs)
	attrs = append(attrs, Attr{"class", class.String()})
	attrs = append(attrs, Attr{"aria-pressed", ternary(ariaPressed, "true", "false")})

	if b.skeleton {
		panic("not implemented")
	}

	if b.href != "" && b.disabled == false {
		if hasIconOnly {
			children := Fragment(
				Raw(`<span class="bx--assistive-text">`),
				Raw(b.iconDescription),
				Raw(`</span>`),
				b.icon,
			)

			render(w, "a", append(attrs, Attr{"href", b.href}), children)
			return
		}

		children := Fragment(
			b.children,
			wrapIcon(b.icon, b.iconDescription),
		)
		render(w, "a", append(attrs, Attr{"href", b.href}), children)
		return
	} else {
		if hasIconOnly {
			children := Fragment(
				Raw(`<span class="bx--assistive-text">`),
				Raw(b.iconDescription),
				Raw(`</span>`),
				b.icon,
			)

			render(w, "button", attrs, children)
			return
		} else {
			children := Fragment(
				b.children,
				wrapIcon(b.icon, b.iconDescription),
			)
			render(w, "button", attrs, children)
			return
		}
	}
}

func wrapIcon(icon Component, description string) Component {
	if icon == nil {
		return nil
	}

	return Fragment(
		Raw(`<div aria-hidden="true" class="bx--btn__icon" style="display: contents;" aria-label="`),
		Raw(description),
		Raw(`">`),
		icon,
		Raw(`</div>`),
	)
}
