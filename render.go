package carbon

import "io"

type Attr struct {
	Name  string
	Value string
}

func renderAttrs(w io.Writer, attrs []Attr) {
	for _, attr := range attrs {
		w.Write([]byte(" "))
		w.Write(yoloBytesUnsafe(attr.Name))
		w.Write([]byte("=\""))
		w.Write(yoloBytesUnsafe(attr.Value))
		w.Write([]byte("\""))
	}
}

func getAttr(attrs []Attr, name string) string {
	for _, attr := range attrs {
		if attr.Name == name {
			return attr.Value
		}
	}

	return ""
}

func renderAll(components []Component, w io.Writer) {
	for _, component := range components {
		component.Render(w)
	}
}
