package chain_dp

import "log"

type Audit interface {
	Audit(days int)
}

type auditor struct {
	name string
	days int
	next Audit
}

func NewAuditor(name string, days int) *auditor {
	return &auditor{name: name, days: days}
}

func (receiver *auditor) SetNext(audit Audit) {
	receiver.next = audit
}

func (a *auditor) Audit(days int) {
	if a.days >= days {
		log.Default().Printf("申请%d天，本次申请由%s审批", days, a.name)
		return
	}
	if a.next != nil {
		a.next.Audit(days)
		return
	}
	log.Default().Printf("申请天数: %d天 不通过", days)
}
