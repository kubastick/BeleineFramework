package beleine

//LABEL

import (
	"errors"
	"fmt"
)

type Label struct {
	id   string
	text string
	size int
	js   string
}

//Pseudo-object creation function
func NewLabel() Label {
	return Label{id: getGlobalID()}
}

func (l *Label) GetID() string {
	return l.id
}

func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) SetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML="%s"`, l.id, text)
}

func (l *Label) SetOnClickListener(listener string) {
	l.js += fmt.Sprintf("%s.onclick = function(){%s}", l.id, listener)
}

///Size in 0-6 number, 0 - default 1-6 <Hx>
func (l *Label) SetSize(size int) error {
	if size > 6 {
		if size < 0 {
			return errors.New("value must be in 0-6 range")
		}
	}
	l.size = size
	return nil
}

func (l *Label) render() string {
	if l.size == 0 {
		return fmt.Sprintf(`<p id="%s">%s</p>`, l.id, l.text)
	} else {
		return fmt.Sprintf(`<h%d id="%s">%s</h%d>`, l.size, l.id, l.text, l.size)
	}
}

func (l *Label) renderJS() string {
	return l.js
}

func (l *Label) GetTextJS() string {
	return fmt.Sprintf("%s.value", l.id)
}
