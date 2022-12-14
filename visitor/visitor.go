package visitor

import "log"

type Visitor interface {
	VisitProgrammer(programmer *Programmer)
	VisitManager(manager *Manager)
}

type CEOVisitor struct {
}

func (receiver *CEOVisitor) VisitProgrammer(programmer *Programmer) {
	log.Default().Println("程序员：", programmer.GetName(), " KPI: ", programmer.GetKpi())
}

func (receiver *CEOVisitor) VisitManager(manager *Manager) {
	log.Default().Println("项目经理：", manager.GetName(), " 项目数：", manager.Product(), " KPI: ", manager.GetKpi())
}

type CTOVisitor struct {
}

func (receiver *CTOVisitor) VisitProgrammer(programmer *Programmer) {
	log.Default().Println("程序员：", programmer.GetName(), " 代码量：", programmer.CodeLine(), " KPI: ", programmer.GetKpi())
}

func (receiver *CTOVisitor) VisitManager(manager *Manager) {
	log.Default().Println("项目经理：", manager.GetName(), " 项目数：", manager.Product(), " KPI: ", manager.GetKpi())

}
