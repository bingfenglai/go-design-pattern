package test

import (
	"lbf.com/go-dp/visitor_dp"
	"log"
	"testing"
)

func TestVisitorDp(t *testing.T) {
	ctoVisitor := &visitor_dp.CTOVisitor{}
	ceoVisitor := &visitor_dp.CEOVisitor{}
	report := visitor_dp.NewReport()
	log.Default().Println("CTO查看报表")
	report.Visit(ctoVisitor)
	log.Default().Println("CEO查看报表")
	report.Visit(ceoVisitor)
}
