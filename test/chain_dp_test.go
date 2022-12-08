package test

import (
	"lbf.com/go-dp/chain_dp"
	"testing"
)

func TestChainDp(t *testing.T) {
	auditor := chain_dp.NewAuditor("组长", 3)
	auditor1 := chain_dp.NewAuditor("部门经理", 5)
	auditor2 := chain_dp.NewAuditor("总经理", 7)
	auditor3 := chain_dp.NewAuditor("BOSS", 10)
	auditor.SetNext(auditor1)
	auditor1.SetNext(auditor2)
	auditor2.SetNext(auditor3)

	days := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, day := range days {
		auditor.Audit(day)
	}
}
