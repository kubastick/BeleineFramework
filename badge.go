package beleine

//BADGE

import (
	"fmt"
	"errors"
)

/*
Badge types:
	0-Primary
	1-Secondary
	2-Success
	3-Danger
	4-Warning
	5-Info
	6-Light
	7-Dark
 */
type Badge struct {
	id string
	text string
	style int
	pill bool
	js string
}

func MakeBadge() Badge {
	return Badge{id:getGlobalID()}
}

func (b *Badge) GetID() string {
	return b.id
}

func (b *Badge) SetText(text string) {
	b.text = text
}

func (b *Badge) SetPill(pill bool) {
	b.pill = pill
}

func (b *Badge) SetStyle(styleNumber int) error {
	if styleNumber>7 { if styleNumber<0 {
		return errors.New("value must be in 0-7 range")
	}}
	b.style = styleNumber
	return nil
}

func (b *Badge) SetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML="%s"`,b.id,text)
}

func (b *Badge) GetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML`,b.id)
}

func (b *Badge) SetOnClickListener(listener string) {
	b.js += fmt.Sprintf("%s.onclick = function(){%s}",b.id,listener)
}

func (b *Badge) render() string {
	pill := ""
	if b.pill {pill="badge-pill "}
	return fmt.Sprintf(`<span id="%s" class="badge %s">%s</span>`,b.id,pill + b.getStyleClass(),b.text)
}

func (b *Badge) renderJS() string {
	return b.js
}

func (b *Badge) getStyleClass() string {
	switch b.style {
	case 0:
		return "badge-primary"
	case 1:
		return "badge-secondary"
	case 2:
		return "badge-success"
	case 3:
		return "badge-danger"
	case 4:
		return "badge-warning"
	case 5:
		return "badge-info"
	case 6:
		return "badge-light"
	case 7:
		return "badge-dark"
	}
	panic("invalid style class number")
}

