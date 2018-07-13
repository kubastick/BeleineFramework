package beleine

//BUTTON

import (
	"errors"
	"fmt"
)

/*
Button types:
	PRIMARY = 0
	SECONDARY = 1
	SUCCES = 2
	DANGER = 3
	WARNING = 4
	IFNO = 5
	LIGHT = 6
	DARK = 7
	LINK = 8
*/

type Button struct {
	id       string
	text     string
	size     string
	btnType  string
	outline  string
	state    bool
	dropdown bool
	ddItems  map[string]Dropdown
	collapse bool
	cItem    Collapse
	js       string
}

//Creates new Button struct
func NewButton() Button {
	return Button{id: GetGlobalID(), btnType: "-primary", state: true}
}

//Returns the button's id
func (b *Button) GetID() string {
	return b.id
}

//Sets text of the button
func (b *Button) SetText(text string) {
	b.text = text
}

//Sets text of the button via JS
func (b *Button) SetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML="%s"`, b.id, text)
}

/*
	Sets size of the button
	Size in 0-2 number, 1 - default
*/
func (b *Button) SetSize(size int) error {
	switch size {
	case 0:
		b.size = "btn-sm"
	case 1:
		b.size = ""
	case 2:
		b.size = "btn-lg"
	default:
		return errors.New("value must be in 0-2 range")
	}
	return nil
}

/*
	Sets type of the button
	Size in 0-8 number, 0 - default (primary)
*/
func (b *Button) SetButtonType(btnType int) error {
	switch btnType {
	case 0:
		b.btnType = "-primary"
	case 1:
		b.btnType = "-secondary"
	case 2:
		b.btnType = "-success"
	case 3:
		b.btnType = "-danger"
	case 4:
		b.btnType = "-warning"
	case 5:
		b.btnType = "-info"
	case 6:
		b.btnType = "-light"
	case 7:
		b.btnType = "-dark"
	case 8:
		b.btnType = "-link"
	default:
		return errors.New("color doesn't exist")
	}
	return nil
}

//Sets outline (default false)
func (b *Button) SetOutline(outline bool) {
	if outline {
		b.outline = "-outline"
	} else {
		b.outline = ""
	}
}

//Sets state of the button (default enabled (true))
func (b *Button) SetState(state bool) {
	b.state = state
}

//Sets JS code function, to be executed after click
func (b *Button) SetOnClickListener(listener string) {
	b.js += fmt.Sprintf("%s.onclick = function(){%s}", b.id, listener)
}

func (b *Button) render() string {
	if b.state {
		if b.dropdown {
			dItems := ""
			for _, v := range b.ddItems {
				dItems += fmt.Sprintf(`<a id="%s" class="dropdown-item" href="#">%s</a>`, v.id, v.name)
			}
			return fmt.Sprintf(`<div class="dropdown">
<button class="btn btn%s%s %s dropdown-toggle" type="button" id="%s" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
%s
</button>
<div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
%s
</div>
</div>`, b.outline, b.btnType, b.size, b.id, b.text, dItems)
		}
		if b.collapse {
			return fmt.Sprintf(`
<p>
  <button id="%s" class="btn btn%s%s %s" type="button" data-toggle="collapse" data-target="#%s" aria-expanded="false" aria-controls="%s">
	%s
  </button>
</p>
<div class="collapse" id="%s">
  <div class="card card-body">
    %s
  </div>
</div>
			`, b.id, b.outline, b.btnType, b.size, b.cItem.id, b.cItem.id, b.text, b.cItem.id, b.cItem.text)
		}
		return fmt.Sprintf(`<button id="%s" type="button" class="btn btn%s%s %s">%s</button>`, b.id, b.outline, b.btnType, b.size, b.text)
	} else {
		return fmt.Sprintf(`<button id="%s" type="button" class="btn btn%s%s %s" disabled>%s</button>`, b.id, b.outline, b.btnType, b.size, b.text)
	}
}

func (b *Button) renderJS() string {
	return b.js
}

func (b *Button) GetTextJS() string {
	return fmt.Sprintf("%s.value", b.id)
}
