package carbon

import (
	"io"
	"strconv"
)

func icon(path []byte) *iconComponent {
	return &iconComponent{
		attrs: nil,
		path:  path,

		size:  16,
		title: "",
	}
}

type iconComponent struct {
	attrs []Attr
	path  []byte

	size  int
	title string
}

func (i *iconComponent) Size(size int) *iconComponent {
	i.size = size
	return i
}

func (i *iconComponent) Title(title string) *iconComponent {
	i.title = title
	return i
}

func (i *iconComponent) Attr(name string, value string) Component {
	i.attrs = append(i.attrs, Attr{name, value})
	return i
}

func (i *iconComponent) Render(w io.Writer) {
	labelled := getAttr(i.attrs, "aria-label")
	if labelled == "" {
		labelled = getAttr(i.attrs, "aria-labelledby")
	}
	if labelled == "" {
		labelled = i.title
	}
	ariaHidden := labelled == ""
	role := ternary(labelled != "", "img", "")

	w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="`))
	w.Write(yoloBytesUnsafe(strconv.Itoa(i.size)))
	w.Write([]byte(`" height="`))
	w.Write(yoloBytesUnsafe(strconv.Itoa(i.size)))
	w.Write([]byte(`"`))
	if ariaHidden {
		w.Write([]byte(` aria-hidden="true"`))
	}
	if role != "" {
		w.Write([]byte(` role="`))
		w.Write(yoloBytesUnsafe(role))
		w.Write([]byte(`"`))
	}
	renderAttrs(w, i.attrs)
	w.Write([]byte(`>`))
	{
		if i.title != "" {
			w.Write([]byte(`<title>`))
			w.Write(yoloBytesUnsafe(i.title))
			w.Write([]byte(`</title>`))
		}
		w.Write(i.path)
	}
	w.Write([]byte(`</svg>`))
}

func EditOff() *iconComponent {
	return icon([]byte(`<path d="M30 28.6L3.4 2 2 3.4l10.1 10.1L4 21.6V28h6.4l8.1-8.1L28.6 30 30 28.6zM9.6 26H6v-3.6l7.5-7.5 3.6 3.6L9.6 26zM29.4 6.2L29.4 6.2l-3.6-3.6c-.8-.8-2-.8-2.8 0l0 0 0 0-8 8 1.4 1.4L20 8.4l3.6 3.6L20 15.6l1.4 1.4 8-8C30.2 8.2 30.2 7 29.4 6.2L29.4 6.2zM25 10.6L21.4 7l3-3L28 7.6 25 10.6z"></path>`))
}

func WarningFilled() *iconComponent {
	return icon([]byte(`<path d="M16,2C8.3,2,2,8.3,2,16s6.3,14,14,14s14-6.3,14-14C30,8.3,23.7,2,16,2z M14.9,8h2.2v11h-2.2V8z M16,25 c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22c0.8,0,1.5,0.7,1.5,1.5S16.8,25,16,25z"></path><path fill="none" d="M17.5,23.5c0,0.8-0.7,1.5-1.5,1.5c-0.8,0-1.5-0.7-1.5-1.5S15.2,22,16,22 C16.8,22,17.5,22.7,17.5,23.5z M17.1,8h-2.2v11h2.2V8z" data-icon-path="inner-path" opacity="0"></path>`))
}

func WarningAltFilled() *iconComponent {
	return icon([]byte(`<path fill="none" d="M16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Zm-1.125-5h2.25V12h-2.25Z" data-icon-path="inner-path"></path><path d="M16.002,6.1714h-.004L4.6487,27.9966,4.6506,28H27.3494l.0019-.0034ZM14.875,12h2.25v9h-2.25ZM16,26a1.5,1.5,0,1,1,1.5-1.5A1.5,1.5,0,0,1,16,26Z"></path><path d="M29,30H3a1,1,0,0,1-.8872-1.4614l13-25a1,1,0,0,1,1.7744,0l13,25A1,1,0,0,1,29,30ZM4.6507,28H27.3493l.002-.0033L16.002,6.1714h-.004L4.6487,27.9967Z"></path>`))
}

func ChevronRight() *iconComponent {
	return icon([]byte(`<path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path>`))
}

func ChevronDown() *iconComponent {
	return icon([]byte(`<path d="M16 22L6 12 7.4 10.6 16 19.2 24.6 10.6 26 12z"></path>`))
}

func Add() *iconComponent {
	return icon([]byte(`<path d="M16 16V8h-2v8H6v2h8v8h2v-8h8v-2z"></path>`))
}

func Delete() *iconComponent {
	return icon([]byte(`<path d="M12 12H14V24H12zM18 12H20V24H18z"></path><path d="M4 6V8H6V28a2 2 0 002 2H24a2 2 0 002-2V8h2V6zM8 28V8H24V28zM12 2H20V4H12z"></path>`))
}
