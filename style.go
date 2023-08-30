package carbon

import (
	_ "embed"
	"io"
	"net/http"

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

var _m *minify.M
var _styleCache map[string]string

func init() {
	_m = minify.New()
	_m.AddFunc("text/css", css.Minify)

	// initialize style cache
	_styleCache = make(map[string]string)
}

var _ Component = (*style)(nil)

type style struct {
	attr  []Attr
	style string
}

func Style(s string) *style {
	s, _ = _m.String("text/css", s)
	return &style{
		attr:  nil,
		style: s,
	}
}

func InlineStyle(href string) *style {
	if s, ok := _styleCache[href]; ok {
		return &style{
			attr:  nil,
			style: s,
		}
	}

	r, err := http.Get(href)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	s, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	minified, _ := _m.String("text/css", string(s))
	_styleCache[href] = minified

	return &style{
		attr:  nil,
		style: minified,
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
