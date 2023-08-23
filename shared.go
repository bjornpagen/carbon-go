package carbon

import (
	"html"
	"io"
	"math/rand"
	"unsafe"

	"github.com/sym01/htmlsanitizer"
)

type Component interface {
	Attr(name string, value string) Component
	Render(w io.Writer)
}

type text struct {
	string
}

func escape(s string) string {
	return html.EscapeString(s)
}

func sanitize(s string) string {
	//TODO: use https://github.com/microcosm-cc/bluemonday
	s, err := htmlsanitizer.SanitizeString(s)
	if err != nil {
		return ""
	}

	return s
}

func Raw(s string) Component {
	return text{s}
}

func Text(s string) Component {
	return text{escape(sanitize(s))}
}

func (t text) Attr(name string, value string) Component {
	return t
}

func (t text) Render(w io.Writer) {
	w.Write(yoloBytesUnsafe(t.string))
}

func yoloBytesUnsafe(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}

	return b
}

var usedIds map[string]bool

func init() {
	usedIds = make(map[string]bool)
}

func useId() string {
	var id string
	for {
		id = randomString(36)
		if usedIds[id] == false {
			usedIds[id] = true
			break
		}
	}
	return id
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type Attr struct {
	Name  string
	Value string
}

func renderAttrs(w io.Writer, attrs []Attr) {
	for _, attr := range attrs {
		w.Write([]byte(` `))
		w.Write(yoloBytesUnsafe(attr.Name))
		w.Write([]byte(`="`))
		w.Write(yoloBytesUnsafe(attr.Value))
		w.Write([]byte(`"`))
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
