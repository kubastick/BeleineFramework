package beleine

import (
	"fmt"
	"errors"
)

//Create By NewNaw()
type Nav struct {
	id string
	style int
	items []NavItem
}

//Types:
//0-default
//1-active
//2-disabled
type NavItem struct {
	Title string
	Link string
	Type int
}

//Adds Nav item
func (n *Nav) AddItem (item *NavItem) error{
	if item.Type < 0 {return errors.New("value outside range")}
	if item.Type > 2 {return errors.New("value outside range")}
	n.items = append(n.items,*item)
	return nil
}

func (n *Nav) render() string {
	result := `<ul class="nav justify-content-center">`
	var classAddon string
	for _,i := range n.items {
		if i.Type == 1 {classAddon="active"}
		if i.Type == 2 {classAddon="disabled"}
		result += fmt.Sprintf(`
<li class="nav-item">
	<a class="nav-link%s" href="%s">%s</a>
</li>`	,classAddon,i.Link,i.Title)
		if i.Type != 0 {classAddon = ""}
	}
	return result
}

func (n *Nav) renderJS() string {
	return ""
	//TODO:Implementation
}