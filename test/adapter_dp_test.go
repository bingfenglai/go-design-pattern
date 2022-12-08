package test

import (
	"lbf.com/go-dp/adapter_dp"
	"testing"
)

func TestAdapterDp(t *testing.T) {
	var adapted adapter_dp.IUserService = &adapter_dp.UserServiceImpl{}

	var target adapter_dp.UserSimpleInfoService = &adapter_dp.UserServiceAdapter{adapted}

	info := target.GetUserSimpleInfo(123)

	t.Log(info)
}
