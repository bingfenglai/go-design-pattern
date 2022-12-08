package compose_dp

type Menu interface {
	Click()
}

// 菜单
type MenuComponent interface {
	Add(component MenuComponent)
	Remove(component MenuComponent)
	GetChild(id int) MenuComponent
	Menu
}
