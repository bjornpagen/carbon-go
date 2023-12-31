package carbon

import (
	"fmt"
	"html"
	"io"
	"math/rand"

	"github.com/sym01/htmlsanitizer"
)

type Component interface {
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

func escape(s string) string {
	return html.EscapeString(s)
}

func sanitize(s string) string {
	//TODO: use https://github.com/microcosm-cc/bluemonday
	r, _ := htmlsanitizer.SanitizeString(s)
	return r
}

func Escaped(s string) string {
	return escape(sanitize(s))
}

func ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}

	return b
}

func useId() string {
	return randomString(36)
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

// TODO: deprecate hack
func renderAttrsWithoutClass(w io.Writer, attrs []Attr) {
	for _, attr := range attrs {
		if attr.Name == "class" {
			continue
		}
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
