package beleine

import "errors"

//COLLAPSE

type Collapse struct {
	id string
	text string
}

func (b *Button) CollapseEnabled(enabled bool) error {
	if enabled && b.dropdown {return errors.New("Dropdown is already initialized!")}
	b.collapse = enabled
	return nil
}

func (b *Button) AddCollapseText(text string) error {
	if !b.collapse {return errors.New("Before you add text to collapse initialize collapse")}
	b.cItem.id = getGlobalID()
	b.cItem.text = text
	return nil
}