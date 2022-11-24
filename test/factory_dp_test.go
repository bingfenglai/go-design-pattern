package test

import (
	"lbf.com/go-dp/factory_dp"
	"testing"
)

func TestFactoryDp(t *testing.T) {

	optionFactory := factory_dp.GetFactoryInstance()

	option, err := optionFactory.CreateOption("?")
	if err != nil {
		t.Error(err)
	}
	if val, err := option.Option(1, 2);err==nil{
		println("计算结果：",val)
	}else {
		println(err.Error())
	}

}
