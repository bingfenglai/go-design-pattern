package iterator_dp

import "github.com/shopspring/decimal"

type MenuItem interface {
	GetName() string
	GetNum() int
	GetPrice() float64
	GetTotal() float64
}

type HdlMenuItem struct {
	name  string
	num   int
	price float64
}

func NewHdlMenuItem(name string, num int, price float64) *HdlMenuItem {
	return &HdlMenuItem{name: name, num: num, price: price}
}

func (receiver *HdlMenuItem) GetName() string {
	return receiver.name
}

func (receiver *HdlMenuItem) GetNum() int {
	return receiver.num
}

func (receiver *HdlMenuItem) GetPrice() float64 {
	return receiver.price
}

func (receiver *HdlMenuItem) GetTotal() float64 {
	f, _ := decimal.NewFromFloat(receiver.price).Mul(decimal.NewFromInt(int64(receiver.num))).Float64()
	return f
}
