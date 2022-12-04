package test

import (
	"lbf.com/go-dp/observer_dp"
	"testing"
)

func TestObserver(t *testing.T) {

	subject := observer_dp.NewUserRegisteredSubject([]observer_dp.Observer{})
	gradeObserver := observer_dp.GradeObserver{}
	subject.Attach(&observer_dp.RightObserver{})
	subject.Attach(&observer_dp.IntegralObserver{})
	subject.Attach(gradeObserver)

	subject.Notify()
	subject.Detach(gradeObserver)
	subject.Notify()

}
