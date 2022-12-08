package compose_dp

import "log"

// 叶子构建
type Leaf struct {
	name string
}

func NewLeaf(name string) MenuComponent {
	return &Leaf{name: name}
}

func (l *Leaf) Add(component MenuComponent) {

}

func (l *Leaf) Remove(component MenuComponent) {

}

func (l *Leaf) GetChild(id int) MenuComponent {
	return nil
}

func (l *Leaf) Click() {
	log.Default().Println("执行菜单操作", l.name)
}
