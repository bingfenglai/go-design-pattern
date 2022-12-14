package visitor

import "math/rand"

type Employee interface {
	GetName() string
	GetKpi() int
	Accept(visitor Visitor)
}

type Programmer struct {
	name     string
	kpi      int
	codeLine int
}

func (receiver *Programmer) CodeLine() int {
	return receiver.codeLine
}

func (receiver *Programmer) GetName() string {
	return receiver.name
}

func (receiver *Programmer) GetKpi() int {
	return receiver.kpi
}

func (receiver *Programmer) Accept(visitor Visitor) {
	visitor.VisitProgrammer(receiver)
}

func NewProgrammer(name string) *Programmer {
	codeLine := rand.Intn(10000)
	kpi := codeLine / 100
	return &Programmer{name: name, kpi: kpi, codeLine: codeLine}
}

type Manager struct {
	name    string
	kpi     int
	product int
}

func (receiver *Manager) Product() int {
	return receiver.product
}

func (receiver *Manager) GetName() string {
	return receiver.name
}

func (receiver *Manager) GetKpi() int {
	return receiver.kpi
}

func (receiver *Manager) Accept(visitor Visitor) {
	visitor.VisitManager(receiver)
}

func NewManager(name string) *Manager {
	product := rand.Intn(10)
	kpi := product * 10
	return &Manager{name: name, kpi: kpi, product: product}
}
