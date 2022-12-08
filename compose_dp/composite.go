package compose_dp

import "log"

type children []MenuComponent

// 树枝构建
type Composite struct {
	children
	name string
}

func (c *Composite) Add(component MenuComponent) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component MenuComponent) {
	if component == nil {
		return
	}
	idx := -1
	for i, child := range c.children {
		if child == component {
			idx = i
			break
		}
	}

	if idx == -1 {
		return
	}

	c.children = append(c.children[:idx], c.children[idx+1:]...)

}

func (c *Composite) GetChild(id int) MenuComponent {

	if id > len(c.children)-1 {
		return nil
	}
	return c.children[id]
}

func (c *Composite) Click() {
	log.Default().Println("执行菜单操作", c.name)
}

func NewComposite(name string) MenuComponent {
	return &Composite{name: name, children: []MenuComponent{}}
}
