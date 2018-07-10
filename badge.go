package beleine

//BADGE

import "fmt"

type Badge struct {
	id string
	text string
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

func (b *Badge) SetTextJS(text string) string {
	return fmt.Sprintf(`%s.innerHTML="%s"`,b.id,text)
}

func (b *Badge) SetOnClickListener(listener string) {
	b.js += fmt.Sprintf("%s.onclick = function(){%s}",b.id,listener)
}

func (b *Badge) render() string {
	return fmt.Sprintf(`<span class="badge badge-secondary">%s</span>`,b.text)
}

func (b *Badge) renderJS() string {
	return b.js
}

