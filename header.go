package carbon

import "io"

type header struct {
	attrs []Attr

	expandedByDefault    bool
	isSideNavOpen        bool
	uiShellAriaLabel     string
	href                 string
	company              any
	platformName         any
	persistHamburgerMenu bool
	expansionBreakpoint  int16
}

var _ Component = (*header)(nil)

func Header() *header {
	return &header{
		attrs: nil,

		expandedByDefault:    false,
		isSideNavOpen:        false,
		uiShellAriaLabel:     "",
		href:                 "",
		company:              nil,
		platformName:         nil,
		persistHamburgerMenu: false,
		expansionBreakpoint:  1056,
	}
}

func (h *header) Attr(name string, value string) *header {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *header) ExpandedByDefault(expandedByDefault bool) *header {
	h.expandedByDefault = expandedByDefault
	return h
}

func (h *header) IsSideNavOpen(isSideNavOpen bool) *header {
	h.isSideNavOpen = isSideNavOpen
	return h
}

func (h *header) UiShellAriaLabel(uiShellAriaLabel string) *header {
	h.uiShellAriaLabel = uiShellAriaLabel
	return h
}

func (h *header) Href(href string) *header {
	h.href = href
	return h
}

func (h *header) Company(company ...any) *header {
	h.company = company
	return h
}

func (h *header) PlatformName(platformName ...any) *header {
	h.platformName = platformName
	return h
}

func (h *header) PersistHamburgerMenu(persistHamburgerMenu bool) *header {
	h.persistHamburgerMenu = persistHamburgerMenu
	return h
}

func (h *header) ExpansionBreakpoint(expansionBreakpoint int16) *header {
	h.expansionBreakpoint = expansionBreakpoint
	return h
}

func (h *header) Render(w io.Writer) {
	w.Write([]byte(`<header class="bx--header" aria-label="`))
	io.WriteString(w, h.uiShellAriaLabel)
	w.Write([]byte(`">`))
	{
		w.Write([]byte(`<a class="bx--header__name" href="`))
		io.WriteString(w, h.href)
		w.Write([]byte(`">`))
		{
			if h.company != nil {
				w.Write([]byte(`<span class="bx--header__name--prefix">`))
				renderAny(w, h.company)
				w.Write([]byte(`&nbsp;</span>`))
			}
			renderAny(w, h.platformName)
		}
		w.Write([]byte(`</a>`))
	}
	w.Write([]byte(`</header>`))
}
