# carbon-go

https://carbon-go.io

every component implements the `carbon.Component` interface:

```go
type Component interface {
	Render(w io.Writer)
	Attr(name, value string) Component
}
```

so, you can use it to write to anything that implements `io.Writer`. use in a server like this:

```go
type foo struct {}

func (f *foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	carbon.Button("Hello world!").Render(w)
}
```

---

currently documentation sucks. just look at the docs for [carbon design svelte](https://carbon-components-svelte.onrender.com/).
i've tried to keep it as similar as possible. moving forward,
we should try to keep upstream with the react version when possible. i just didn't want to comb through react code.


will be adding doc comments.