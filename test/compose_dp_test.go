package test

import (
	"lbf.com/go-dp/compose_dp"
	"testing"
)

func TestComposeDp(t *testing.T) {
	rootMenu := compose_dp.NewComposite("根菜单")
	authorityMenu := compose_dp.NewComposite("权限管理")
	roleMenu := compose_dp.NewLeaf("角色管理")
	authorityMenu.Add(roleMenu)
	monitorMenu := compose_dp.NewComposite("系统监控")
	sessionMenu := compose_dp.NewLeaf("在线会话")
	monitorMenu.Add(sessionMenu)
	rootMenu.Add(authorityMenu)
	rootMenu.Add(monitorMenu)

	rootMenu.Click()
	rootMenu.GetChild(1).GetChild(0).Click()

}
