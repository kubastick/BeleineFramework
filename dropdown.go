package beleine

//DROPDOWN

type Dropdown struct {
	id string
	name string
}

func (b *Button) DropdownEnabled(enable bool) {
	b.dropdown = enable
	if enable {
		b.ddItems = make(map[string]Dropdown)
	}
}

func (b *Button) AddDropdownItem(name string) {
	b.ddItems[name] = Dropdown{id:getGlobalID(),name:name}
}

func (b *Button) GetDropdownItemID(itemName string) string {
	return b.ddItems[itemName].id
}
