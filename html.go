package carbon

import "io"

type html_ struct {
	attrs []Attr

	head any
	body any
}

var _ Component = (*html_)(nil)

func Html() *html_ {
	return &html_{
		attrs: nil,

		head: nil,
		body: nil,
	}
}

func (h *html_) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *html_) Head(head any) *html_ {
	h.head = head
	return h
}

func (h *html_) Body(body any) *html_ {
	h.body = body
	return h
}

func (h *html_) Render(w io.Writer) {
	w.Write([]byte("<!DOCTYPE html>"))
	w.Write([]byte("<html"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	w.Write([]byte("<head>"))
	renderAny(w, h.head)
	w.Write([]byte("</head>"))
	w.Write([]byte("<body>"))
	renderAny(w, h.body)
	w.Write([]byte("</body>"))
	w.Write([]byte("</html>"))
}
