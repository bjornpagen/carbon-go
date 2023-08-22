package carbon

// import (
// 	"io"
// )

// type fragment struct {
// 	children []Component
// }

// var _ Component = (*fragment)(nil)

// func (f *fragment) Attr(name string, value string) Component {
// 	return f
// }

// func (f *fragment) Render(w io.Writer) {
// 	for _, c := range f.children {
// 		if c != nil {
// 			c.Render(w)
// 		}
// 	}
// }

// func Fragment(children ...Component) *fragment {
// 	return &fragment{
// 		children: children,
// 	}
// }
