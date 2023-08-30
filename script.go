package carbon

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

var _scriptCache map[string]string

func init() {
	_m = minify.New()
	_m.AddFunc("text/javascript", js.Minify)

	// initialize script cache
	_scriptCache = make(map[string]string)
}

var _ Component = (*script)(nil)

type script struct {
	attr   []Attr
	script string
}

func Script(s string) *script {
	s, _ = _m.String("text/javascript", s)
	return &script{
		attr:   nil,
		script: s,
	}
}

func InlineScript(href string) *script {
	if s, ok := _scriptCache[href]; ok {
		return &script{
			attr:   nil,
			script: s,
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
	_scriptCache[href] = minified

	return &script{
		attr:   nil,
		script: minified,
	}
}

func (s *script) Attr(name string, value string) Component {
	s.attr = append(s.attr, Attr{name, value})
	return s
}

func (s *script) Render(w io.Writer) {
	w.Write([]byte(`<script`))
	renderAttrs(w, s.attr)
	w.Write([]byte(`>`))
	{
		io.WriteString(w, s.script)
	}
	w.Write([]byte(`</script>`))
}
