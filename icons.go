package carbon

import (
	"io"
	"strconv"
)

// TODO: autogenerate these

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
	io.WriteString(w, strconv.Itoa(i.size))
	w.Write([]byte(`" height="`))
	io.WriteString(w, strconv.Itoa(i.size))
	w.Write([]byte(`"`))
	if ariaHidden {
		w.Write([]byte(` aria-hidden="true"`))
	}
	if role != "" {
		w.Write([]byte(` role="`))
		io.WriteString(w, role)
		w.Write([]byte(`"`))
	}
	renderAttrs(w, i.attrs)
	w.Write([]byte(`>`))
	{
		if i.title != "" {
			w.Write([]byte(`<title>`))
			io.WriteString(w, i.title)
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

func Checkmark() *iconComponent {
	return icon([]byte(`<path d="M13 24L4 15 5.414 13.586 13 21.171 26.586 7.586 28 9 13 24z"></path>`))
}

func GitHub() *iconComponent {
	return icon([]byte(`<path fill-rule="evenodd" d="M16,2a14,14,0,0,0-4.43,27.28c.7.13,1-.3,1-.67s0-1.21,0-2.38c-3.89.84-4.71-1.88-4.71-1.88A3.71,3.71,0,0,0,6.24,22.3c-1.27-.86.1-.85.1-.85A2.94,2.94,0,0,1,8.48,22.9a3,3,0,0,0,4.08,1.16,2.93,2.93,0,0,1,.88-1.87c-3.1-.36-6.37-1.56-6.37-6.92a5.4,5.4,0,0,1,1.44-3.76,5,5,0,0,1,.14-3.7s1.17-.38,3.85,1.43a13.3,13.3,0,0,1,7,0c2.67-1.81,3.84-1.43,3.84-1.43a5,5,0,0,1,.14,3.7,5.4,5.4,0,0,1,1.44,3.76c0,5.38-3.27,6.56-6.39,6.91a3.33,3.33,0,0,1,.95,2.59c0,1.87,0,3.38,0,3.84s.25.81,1,.67A14,14,0,0,0,16,2Z"></path>`))
}

func Copy() *iconComponent {
	return icon([]byte(`<path d="M28,10V28H10V10H28m0-2H10a2,2,0,0,0-2,2V28a2,2,0,0,0,2,2H28a2,2,0,0,0,2-2V10a2,2,0,0,0-2-2Z"></path><path d="M4,18H2V4A2,2,0,0,1,4,2H18V4H4Z"></path>`))
}

func View() *iconComponent {
	return icon([]byte(`<path d="M30.94,15.66A16.69,16.69,0,0,0,16,5,16.69,16.69,0,0,0,1.06,15.66a1,1,0,0,0,0,.68A16.69,16.69,0,0,0,16,27,16.69,16.69,0,0,0,30.94,16.34,1,1,0,0,0,30.94,15.66ZM16,25c-5.3,0-10.9-3.93-12.93-9C5.1,10.93,10.7,7,16,7s10.9,3.93,12.93,9C26.9,21.07,21.3,25,16,25Z"></path><path d="M16,10a6,6,0,1,0,6,6A6,6,0,0,0,16,10Zm0,10a4,4,0,1,1,4-4A4,4,0,0,1,16,20Z"></path>`))
}

func ViewOff() *iconComponent {
	return icon([]byte(`<path d="M5.24,22.51l1.43-1.42A14.06,14.06,0,0,1,3.07,16C5.1,10.93,10.7,7,16,7a12.38,12.38,0,0,1,4,.72l1.55-1.56A14.72,14.72,0,0,0,16,5,16.69,16.69,0,0,0,1.06,15.66a1,1,0,0,0,0,.68A16,16,0,0,0,5.24,22.51Z"></path><path d="M12 15.73a4 4 0 013.7-3.7l1.81-1.82a6 6 0 00-7.33 7.33zM30.94 15.66A16.4 16.4 0 0025.2 8.22L30 3.41 28.59 2 2 28.59 3.41 30l5.1-5.1A15.29 15.29 0 0016 27 16.69 16.69 0 0030.94 16.34 1 1 0 0030.94 15.66zM20 16a4 4 0 01-6 3.44L19.44 14A4 4 0 0120 16zm-4 9a13.05 13.05 0 01-6-1.58l2.54-2.54a6 6 0 008.35-8.35l2.87-2.87A14.54 14.54 0 0128.93 16C26.9 21.07 21.3 25 16 25z"></path>`))
}
