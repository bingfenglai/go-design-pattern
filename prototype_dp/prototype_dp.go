package prototype_dp

type Cloneable interface {
	Clone() Cloneable
}

type Immortal struct {
	Name string
	Age  int
}

func (receiver *Immortal) Clone() Cloneable {
	return &Immortal{
		Name: receiver.Name,
		Age:  receiver.Age,
	}
}

func CreateImmortal(name string, age int) *Immortal {

	return &Immortal{
		Name: name,
		Age:  age,
	}

}
