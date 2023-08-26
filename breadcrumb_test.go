package carbon

import (
	"bytes"
	"testing"
)

func TestBreadCrumb(t *testing.T) {
	expected := ``
	b := bytes.Buffer{}
	Breadcrumb(
		BreadcrumbItem("Dashboard").Href("/"),
		BreadcrumbItem("Annual reports").Href("/reports"),
		BreadcrumbItem("2019").Href("/reports/2019").IsCurrentPage(true),
	).Render(&b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}
