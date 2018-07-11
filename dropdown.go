package beleine

import "errors"

//DROPDOWN

type Dropdown struct {
	id string
	name string
}

func (b *Button) DropdownEnabled(enable bool) error {
	if enable && b.collapse {return errors.New("Collapse is already initialized!")}
	b.dropdown = enable
	if enable {
		b.ddItems = make(map[string]Dropdown)
	}
	return nil
}

func (b *Button) AddDropdownItem(name string) error {
	if (!b.dropdown) {return errors.New("Before you add item to dropdown initialize dropdown")}
	b.ddItems[name] = Dropdown{id:getGlobalID(),name:name}
	return nil
}

func (b *Button) GetDropdownItemID(itemName string) string {
	return b.ddItems[itemName].id
}
