package test

import (
	"lbf.com/go-dp/decorator_dp"
	"testing"
)

func TestDecoratorDp(t *testing.T) {

	api := decorator_dp.GetResourceApi{}

	decoratorApi := &decorator_dp.EncryptionDecorator{Decorator: &api}
	//decoratorApi.Access()
	// 二次装饰
	decorator := &decorator_dp.SecurityDecorator{Decorator: decoratorApi}

	decorator.Access()
}
