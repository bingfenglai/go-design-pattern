package test

import (
	"lbf.com/go-dp/factory_method_dp"
	"testing"
)

func TestFm(t *testing.T) {

	x, y := 9.9, 5.0

	var factory factory_method_dp.Factory

	factory = &factory_method_dp.OperationAddFactory{}

	result := factory.CreateOperation().GetResult(x, y)
	t.Log("入参：", x, ",", y)
	t.Log("加法运算结果: ", result)

	factory = &factory_method_dp.OperationSubFactory{}
	result = factory.CreateOperation().GetResult(x, y)
	t.Log("减法运算结果: ", result)

	factory = &factory_method_dp.OperationMulFactory{}
	result = factory.CreateOperation().GetResult(x, y)
	t.Log("乘法运算结果: ", result)

	factory = &factory_method_dp.OperationDivFactory{}
	result = factory.CreateOperation().GetResult(x, y)
	t.Log("除法运算结果: ", result)

}
