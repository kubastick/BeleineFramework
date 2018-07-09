package main

import "fmt"

type Label struct {
	id string
	text string
}

func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) render() string {
	return fmt.Sprintf(`<p>%s</p>`,l.text)
}
