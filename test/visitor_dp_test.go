package test

import (
	"lbf.com/go-dp/visitor"
	"log"
	"testing"
)

func TestVisitorDp(t *testing.T) {
	ctoVisitor := &visitor.CTOVisitor{}
	ceoVisitor := &visitor.CEOVisitor{}
	report := visitor.NewReport()
	log.Default().Println("CTO查看报表")
	report.Visit(ctoVisitor)
	log.Default().Println("CEO查看报表")
	report.Visit(ceoVisitor)
}
