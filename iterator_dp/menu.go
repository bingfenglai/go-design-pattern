package iterator_dp

type Menu interface {
	Iterator() Iterator
}

type HdlMenu struct {
	menuItem []MenuItem
}

func (menu *HdlMenu) Iterator() Iterator {
	return NewHdlMenuIterator(menu.menuItem)
}

func NewHdlMenu() *HdlMenu {
	return &HdlMenu{menuItem: []MenuItem{}}
}

func (menu *HdlMenu) AddItem(item MenuItem) *HdlMenu {
	menu.menuItem = append(menu.menuItem, item)
	return menu
}
