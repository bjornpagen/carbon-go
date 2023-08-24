package carbon

import (
	"fmt"
	"html"
	"io"
	"math/rand"

	"github.com/sym01/htmlsanitizer"
)

type Component interface {
	Attr(name string, value string) Component
	Render(w io.Writer)
}

func renderAny(w io.Writer, a any) {
	switch a := a.(type) {
	case []any:
		for _, a := range a {
			renderAny(w, a)
		}
	case interface{ Render(io.Writer) }:
		a.Render(w)
	case func(io.Writer):
		a(w)
	case []byte:
		w.Write(a)
	case string:
		io.WriteString(w, a)
	case fmt.Stringer:
		io.WriteString(w, a.String())
	case nil:
		// ignore
	default:
		panic(fmt.Sprintf("unsupported type %T", a))
	}
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

func Sanitized(s string) Component {
	return text{escape(sanitize(s))}
}

func (t text) Attr(name string, value string) Component {
	return t
}

func (t text) Render(w io.Writer) {
	io.WriteString(w, t.string)
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
		io.WriteString(w, attr.Name)
		w.Write([]byte(`="`))
		io.WriteString(w, attr.Value)
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
