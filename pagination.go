package beleine

//PAGINATION

import (
	"errors"
	"fmt"
)

type Pagination struct {
	id         string
	size       string
	alignment  string
	components []Component
	js         string
}

//Pseudo-object creation function
func NewPagination() Pagination {
	return Pagination{id: GetGlobalID()}
}

func (p *Pagination) GetID() string {
	return p.id
}

func (p *Pagination) SetAlignment(alignment string) error {
	switch alignment {
	case "NONE":
		p.alignment = ""
	case "CENTER":
		p.alignment = "justify-content-center"
	case "END":
		p.alignment = "justify-content-end"
	default:
		return errors.New("there is no such alignment")
	}
	return nil
}

///Size in 0-6 number, 0 - default 1-6 <Hx>
func (p *Pagination) SetSize(size int) error {
	switch size {
	case 0:
		p.size = "pagination-sm"
	case 1:
		p.size = ""
	case 2:
		p.size = "pagination-lg"
	default:
		return errors.New("value must be in 0-2 range")
	}
	return nil
}

func (p *Pagination) AddItem(c Component) {
	p.components = append(p.components, c)
}

func (p *Pagination) Render() string {
	items := ""

	for _, v := range p.components {
		items += v.Render()
	}

	return fmt.Sprintf(`
<nav id="%s" aria-label="Page navigation example">
  <ul class="pagination %s %s">
	%s
  </ul>
</nav>
`, p.id, p.size, p.alignment, items)
}

func (p *Pagination) RenderJS() string {
	return p.js
}
