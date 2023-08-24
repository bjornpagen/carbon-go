package carbon

import (
	"fmt"
	"io"

	"github.com/valyala/bytebufferpool"
	"golang.org/x/exp/slices"
)

type button struct {
	attrs    []Attr
	children any

	kind             string
	size             string
	expressive       bool
	isSelected       bool
	icon             Component
	iconDescription  string
	tooltipAlignment string
	tooltipPosition  string
	disabled         bool
	href             string
	type_            string
}

var _ Component = (*button)(nil)

func Button(children ...any) *button {
	return &button{
		children: children,
		attrs:    nil,

		kind:             "primary",
		size:             "",
		expressive:       false,
		isSelected:       false,
		icon:             nil,
		iconDescription:  "",
		tooltipAlignment: "center",
		tooltipPosition:  "bottom",
		disabled:         false,
		href:             "",
		type_:            "button",
	}
}

func (b *button) Attr(name string, value string) Component {
	b.attrs = append(b.attrs, Attr{name, value})
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
	var sizes = []string{"", "field", "sm", "lg", "xl"}
	if slices.Contains(sizes, size) == false {
		panic("invalid size, options are: " + fmt.Sprintf("%v", sizes))
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
	if hasIconOnly && b.iconDescription == "" {
		panic("iconDescription is required when icon is specified")
	}
	ariaPressed := ternary(hasIconOnly && b.kind == "ghost" && b.href == "", b.isSelected, false)
	renderIconInButton := renderIconInButtonClosure(b.icon, b.iconDescription)

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

	attrs := append(b.attrs, Attr{"class", class.String()}, Attr{"aria-pressed", ternary(ariaPressed, "true", "false")})

	if b.href != "" && b.disabled == false && hasIconOnly {
		w.Write([]byte("<a"))
		renderAttrs(w, attrs)
		w.Write([]byte(" href=\""))
		w.Write([]byte(b.href))
		w.Write([]byte("\">"))
		w.Write([]byte("<span class=\"bx--assistive-text\">"))
		io.WriteString(w, b.iconDescription)
		w.Write([]byte("</span>"))
		b.icon.Attr("class", "bx--btn__icon").Attr("aria-hidden", "true").Attr("aria-label", b.iconDescription).Render(w)
		w.Write([]byte("</a>"))

		return
	}

	if b.href != "" && b.disabled == false {
		w.Write([]byte("<a"))
		renderAttrs(w, attrs)
		w.Write([]byte(" href=\""))
		w.Write([]byte(b.href))
		w.Write([]byte("\">"))
		renderAny(w, b.children)
		renderIconInButton(w)
		w.Write([]byte("</a>"))

		return
	}

	if hasIconOnly {
		w.Write([]byte("<button"))
		renderAttrs(w, attrs)
		w.Write([]byte(">"))
		w.Write([]byte("<span class=\"bx--assistive-text\">"))
		io.WriteString(w, b.iconDescription)
		w.Write([]byte("</span>"))
		b.icon.Attr("class", "bx--btn__icon").Attr("aria-hidden", "true").Attr("aria-label", b.iconDescription).Render(w)
		w.Write([]byte("</button>"))

		return
	}

	w.Write([]byte("<button"))
	renderAttrs(w, attrs)
	w.Write([]byte(">"))
	renderAny(w, b.children)
	renderIconInButton(w)
	w.Write([]byte("</button>"))
}

func renderIconInButtonClosure(icon Component, description string) func(w io.Writer) {
	if icon == nil {
		return func(w io.Writer) {}
	}

	return func(w io.Writer) {
		icon.Attr("class", "bx--btn__icon").Attr("aria-hidden", "true").Attr("aria-label", description).Render(w)
	}
}
