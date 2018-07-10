package main

import (
	"fmt"
	"strconv"
	"errors"
)

type Label struct {
	id string
	text string
	size int
}

//Pseudo-object creation function
func MakeLabel() Label {
	return Label{id:strconv.Itoa(getGlobalID())}
}

func (l *Label) GetLabelId() string {
	return l.id
}

func (l *Label) SetText(text string) {
	l.text = text
}

///Size in 0-6 number, 0 - default 1-6 <Hx>
func (l *Label) SetSize(size int) error {
	if size>6 { if size<0 {
		return errors.New("value must be in 0-6 range")
	}}
	l.size = size
	return nil
}

func (l *Label) render() string {
	if l.size==0 {
		return fmt.Sprintf(`<p id="%s">%s</p>`,l.id,l.text)
	} else {
		return fmt.Sprintf(`<h%d id="%s">%s</h%d>`,l.size,l.id,l.text,l.size)
	}
}
