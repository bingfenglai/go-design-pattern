package test

import (
	"lbf.com/go-dp/proxy_dp"
	"testing"
)

func TestProxy(t *testing.T) {

	service := proxy_dp.CreateResourceService()
	if resource, err := service.GetResource("9527"); err == nil {
		t.Log(resource)
	} else {
		t.Error(err.Error())
	}
}
