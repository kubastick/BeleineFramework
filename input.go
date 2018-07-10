package beleine

//INPUT

import (
	"fmt"
	"errors"
)


type Input struct {
	id string
	text string
	hint string
	size string
	inputType string
	js string
}


//Pseudo-object creation function
func MakeInput() Input {
	return Input{id:getGlobalID(), inputType:"text"}
}

func (i *Input) GetID() string {
	return i.id
}

func (i *Input) SetTextJS(text string) string {
	i.text = text
	return fmt.Sprintf(`%s.value="%s"`,i.id,i.text)
}

func (i *Input) SetHint(hint string) {
	i.hint = hint
}

///Size in 0-2 number, 1 - default
func (i *Input) SetSize(size int) error {
	if size>2 { if size<0 {
		return errors.New("value must be in 0-2 range")
	}}

	switch size {
	case 0:
		i.size = "form-control-sm"
	case 1:
		i.size = ""
	case 2:
		i.size = "form-control-lg"
	}
	return nil
}

func (i *Input) SetInputType(inputType string) {
	i.inputType = inputType
}

//WIP
func (i *Input) GetValueJS() {
	i.js += fmt.Sprintf("console.log(%s.value);", i.id)
}

func (i *Input) render() string {
	return fmt.Sprintf(`<input id="%s" type="%s" class="form-control %s" placeholder="%s">`, i.id, i.inputType, i.size, i.hint)
}

func (i *Input) renderJS() string {
	return i.js
}

