package main

import (
	"fmt"
	"strconv"
	"errors"
)

/*
Button types:
	PRIMARY = 0
	SUCCES = 1
	IFNO = 2
	WARNING = 3
	DANGER = 4
	PRIMARY = 5
	SECONDARY = 6
	LIGHT = 7
	DARK = 8
*/

type Button struct {
	id string
	text string
	size string
	btnType string
	outline string
	state bool
}


//Pseudo-object creation function
func MakeButton() Button {
	return Button{id:strconv.Itoa(getGlobalID()),btnType:"-primary",state:true}
}

func (b *Button) SetText(text string) {
	b.text = text
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


func (b *Button) render() string {


	if b.state {
		return fmt.Sprintf(`<button type="button" class="btn btn%s%s %s">%s</button>`, b.outline, b.btnType, b.size, b.text)
	} else {
		return fmt.Sprintf(`<button type="button" class="btn btn%s%s %s" disabled>%s</button>`, b.outline, b.btnType, b.size, b.text)
	}

}
