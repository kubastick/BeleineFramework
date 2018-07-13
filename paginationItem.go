package beleine

import "fmt"

//PAGINATION ITEM

type PaginationItem struct {
	id      string
	text    string
	enabled bool
	js      string
}

func NewPaginationItem() PaginationItem {
	return PaginationItem{id: GetGlobalID(), enabled: true}
}

func (p *PaginationItem) SetText(text string) {
	p.text = text
}

func (p *PaginationItem) GetID() string {
	return p.id
}

func (p *PaginationItem) SetEnabled(enabled bool) {
	p.enabled = enabled
}

func (p *PaginationItem) SetOnClickListener(listener string) {
	p.js += fmt.Sprintf("%s.onclick = function(){%s}", p.id, listener)
}

func (p *PaginationItem) render() string {
	if p.enabled {
		return fmt.Sprintf(`<li id="%s" class="page-item"><a class="page-link" href="#">%s</a></li>`, p.id, p.text)
	} else {
		return fmt.Sprintf(`<li id="%s" class="page-item disabled"><a class="page-link" href="#">%s</a></li>`, p.id, p.text)
	}
}

func (p *PaginationItem) renderJS() string {
	return p.js
}
