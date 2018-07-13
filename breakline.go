package beleine

type Breakline struct {
	id string
}

func NewBreakline() Breakline {
	return Breakline{id: getGlobalID()}
}

func (b *Breakline) render() string {
	return "</br>"
}

func (b *Breakline) renderJS() string {
	return ""
}

func (b *Breakline) GetID() string {
	return b.id
}
