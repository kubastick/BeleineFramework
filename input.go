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


//Create new Input struct
func NewInput() Input {
	return Input{id:getGlobalID(), inputType:"text"}
}

//Returns ID of component
func (i *Input) GetID() string {
	return i.id
}

//Return JS code as string, that sets Input text
func (i *Input) SetTextJS(text string) string {
	i.text = text
	return fmt.Sprintf(`%s.value="%s"`,i.id,i.text)
}

//Sets input hint
func (i *Input) SetHint(hint string) {
	i.hint = hint
}

//Size in 0-2 number, 1 - default
//Throws error when size is not in 0-2 range
func (i *Input) SetSize(size int) error {
	switch size {
	case 0:
		i.size = "form-control-sm"
	case 1:
		i.size = ""
	case 2:
		i.size = "form-control-lg"
	default:
		return errors.New("value must be in 0-2 range")
	}
	return nil
}

//Sets input type
//TODO:Make switch instead of html code
func (i *Input) SetInputType(inputType string) {
	i.inputType = inputType
}

//Return JS code as string, that returns value of input
func (i *Input) GetTextJS() string {
	return fmt.Sprintf("%s.value", i.id)
}

//Sets js code, to be executed after click
func (i *Input) SetOnClickListener(listener string)  {
	i.js += fmt.Sprintf("%s.onclick = function(){%s}",i.id,listener)
}

func (i *Input) render() string {
	return fmt.Sprintf(`<input id="%s" type="%s" class="form-control %s" placeholder="%s">`, i.id, i.inputType, i.size, i.hint)
}

func (i *Input) renderJS() string {
	return i.js
}

