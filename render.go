package carbon

import "io"

type Attr struct {
	Name  string
	Value string
}

func renderAttrs(w io.Writer, attrs []Attr) {
	for _, attr := range attrs {
		w.Write([]byte(" "))
		w.Write([]byte(attr.Name))
		w.Write([]byte("=\""))
		w.Write([]byte(attr.Value))
		w.Write([]byte("\""))
	}
}

func render(w io.Writer, tag string, attrs []Attr, children Component) {
	io.WriteString(w, "<")
	io.WriteString(w, tag)
	renderAttrs(w, attrs)
	io.WriteString(w, ">")
	children.Render(w)
	io.WriteString(w, "</")
	io.WriteString(w, tag)
	io.WriteString(w, ">")
}
