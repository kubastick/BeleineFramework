package beleine

import (
	"errors"
	"fmt"
)

//Create By NewNaw()
type Nav struct {
	id      string
	style   int
	justify int
	expand  bool
	items   []NavItem
}

//Types:
//0-default
//1-active
//2-disabled
type NavItem struct {
	Title string
	Link  string
	Type  int
}

func NewNav() Nav {
	return Nav{id: getGlobalID()}
}

//Adds Nav item
func (n *Nav) AddItem(item *NavItem) error {
	if item.Type < 0 {
		return errors.New("value outside range")
	}
	if item.Type > 2 {
		return errors.New("value outside range")
	}
	n.items = append(n.items, *item)
	return nil
}

//Arguments:
//0-default justify
//1-center horizontal
//2-push to left
//3-vertical
func (n *Nav) SetJustify(justify int) error {
	if justify < 0 {
		return errors.New("value outside range")
	}
	if justify > 3 {
		return errors.New("value outside range")
	}
	n.justify = justify
	return nil
}

//Arguments:
//0-default
//1-tabs
//2-pills
func (n *Nav) SetStyle(style int) error {
	if style < 0 {
		return errors.New("value outside range")
	}
	if style > 2 {
		return errors.New("value outside range")
	}
	n.style = style
	return nil
}

//Enable/disable expanding
func (n *Nav) SetExpand(value bool) {
	n.expand = value
}

func (n *Nav) render() string {
	result := fmt.Sprintf(`<ul class="nav%s%s%s">`, n.getStyleTag(), n.getJustifyTag(), n.getExpandTag())
	var classAddon string
	for _, i := range n.items {
		if i.Type == 1 {
			classAddon = "active"
		}
		if i.Type == 2 {
			classAddon = "disabled"
		}
		result += fmt.Sprintf(`
<li class="nav-item">
	<a class="nav-link%s" href="%s">%s</a>
</li>
`, classAddon, i.Link, i.Title)
		if i.Type != 0 {
			classAddon = ""
		}
	}
	result += `</ul>`
	return result
}

func (n *Nav) renderJS() string {
	return ""
	//TODO:Implementation
}

func (n *Nav) getStyleTag() string {
	switch n.style {
	case 0:
		return ""
	case 1:
		return " nav-tabs"
	case 2:
		return " nav-pills"
	}
	panic("invalid style value")
}

func (n *Nav) getExpandTag() string {
	if n.expand {
		return " nav-fill"
	}
	return ""
}

func (n *Nav) getJustifyTag() string {
	switch n.justify {
	case 0:
		return ""
	case 1:
		return " justify-content-center"
	case 2:
		return " justify-content-end"
	case 3:
		return " flex-column"
	}
	panic("invalid justify value")
}

func (n *Nav) GetID() string {
	return n.id
}
