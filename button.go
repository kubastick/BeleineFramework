package beleine

//BUTTON

import (
	"fmt"
	"errors"
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
	id string
	text string
	size string
	btnType string
	outline string
	state bool
	dropdown bool
	ddItems map[string]dropdownItem
	js string
}

type dropdownItem struct {
	id string
	name string
}


//Pseudo-object creation function
func MakeButton() Button {
	return Button{id:getGlobalID(),btnType:"-primary",state:true}
}

func (b *Button) DropdownEnabled(enable bool) {
	b.dropdown = enable
	if enable {
		b.ddItems = make(map[string]dropdownItem)
	}
}

func (b *Button) AddDropdownItem(name string) string {
	b.ddItems[name] = dropdownItem{name:name}
	//b.ddItems[name].id = getGlobalID()
	return b.ddItems["dTest"].name
}

func (b *Button) GetID() string {
	return b.id
}

func (b *Button) SetText(text string) {
	b.text = text
}

func (b *Button) SetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML="%s"`,b.id,text)
}

///Size in 0-2 number, 1 - default
func (b *Button) SetSize(size int) error {
	if size>2 { if size<0 {
		return errors.New("value must be in 0-2 range")
	}}

	switch size {
	case 0:
		b.size = "btn-sm"
	case 1:
		b.size = ""
	case 2:
		b.size = "btn-lg"
	}
	return nil
}

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

func (b *Button) SetOutline(outline bool) {
	if outline {
		b.outline = "-outline"
	} else {
		b.outline = ""
	}
}

func (b *Button) SetState(state bool) {
		b.state = state
}

func (b *Button) SetOnClickListener(listener string)  {
	b.js += fmt.Sprintf("%s.onclick = function(){%s}",b.id,listener)
}

func (b *Button) render() string {
	if b.state {
		return fmt.Sprintf(`<button id="%s" type="button" class="btn btn%s%s %s">%s</button>`, b.id, b.outline, b.btnType, b.size, b.text)
	} else {
		return fmt.Sprintf(`<button id="%s" type="button" class="btn btn%s%s %s" disabled>%s</button>`, b.id, b.outline, b.btnType, b.size, b.text)
	}
}

func (b *Button) renderJS() string {
	return b.js
}

func (b *Button) GetTextJS() string{
	return fmt.Sprintf("%s.value", b.id)
}
