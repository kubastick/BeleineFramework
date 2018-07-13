package beleine

type Breakline struct {
	id string
}

func NewBreakline() Breakline {
	return Breakline{id: GetGlobalID()}
}

func (b *Breakline) Render() string {
	return "</br>"
}

func (b *Breakline) RenderJS() string {
	return ""
}

func (b *Breakline) GetID() string {
	return b.id
}
