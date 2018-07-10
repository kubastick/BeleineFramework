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

func (l *Label) SetText(text string) string {
	l.text = text
	return fmt.Sprintf("%s.innerHTML=%s",l.id,text)
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
		return fmt.Sprintf(`<p>%s</p>`, l.text)
	} else {
		return fmt.Sprintf(`<h%d>%s</h%d>`,l.size,l.text,l.size)
	}
}
