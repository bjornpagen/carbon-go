package carbon

import (
	_ "embed"
	"io"

	"github.com/efficientgo/core/merrors"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

//go:embed css/base.css
var baseCss string

//go:embed css/font-family.css
var fontFamily string

func init() {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	errs := merrors.New()
	var err error
	baseCss, err = m.String("text/css", baseCss)
	errs.Add(err)
	fontFamily, err = m.String("text/css", fontFamily)
	errs.Add(err)
	if errs.Err() != nil {
		panic(errs.Err())
	}
}

var _ Component = (*style)(nil)

type style struct {
	attr   []Attr
	styles []string
}

func Style(styles ...string) *style {
	return &style{
		styles: styles,
	}
}

func (s *style) Attr(name string, value string) Component {
	s.attr = append(s.attr, Attr{name, value})
	return s
}

func (s *style) DefaultFonts() *style {
	s.styles = append(s.styles, fontFamily)
	return s
}

func (s *style) Render(w io.Writer) {
	w.Write([]byte(`<style>`))
	w.Write([]byte(baseCss))
	for _, style := range s.styles {
		w.Write([]byte(style))
	}
	w.Write([]byte(`</style>`))
}
