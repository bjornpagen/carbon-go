package carbon

import (
	_ "embed"
	"io"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

func init() {
	BaseCss += baseGoCss
}

//go:embed css/base.css
var BaseCss string

//go:embed css/base-go.css
var baseGoCss string

//go:embed css/font-family.css
var BaseFontCss string

var m *minify.M

func init() {
	m = minify.New()
	m.AddFunc("text/css", css.Minify)
}

var _ Component = (*style)(nil)

type style struct {
	attr  []Attr
	style string
}

func Style(s string) *style {
	s, _ = m.String("text/css", s)
	return &style{
		attr:  nil,
		style: s,
	}
}

func (s *style) Attr(name string, value string) Component {
	s.attr = append(s.attr, Attr{name, value})
	return s
}

func (s *style) Render(w io.Writer) {
	w.Write([]byte(`<style`))
	renderAttrs(w, s.attr)
	w.Write([]byte(`>`))
	{
		io.WriteString(w, s.style)
	}
	w.Write([]byte(`</style>`))
}
