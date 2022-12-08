package test

import (
	"lbf.com/go-dp/singleton_dp"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := singleton_dp.GetInstance()
	instance2 := singleton_dp.GetInstance()
	t.Log(instance1 == instance2)

	for i := 0; i < 100; i++ {
		go singleton_dp.GetInstanceV2()
	}

}
