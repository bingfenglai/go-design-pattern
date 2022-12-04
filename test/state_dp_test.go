package test

import (
	"lbf.com/go-dp/state_dp"
	"testing"
)

func TestStateDp(t *testing.T) {
	var ctx *state_dp.Context

	ctx = state_dp.NewContext(state_dp.NewUnMemberState())
	ctx.Request()
	ctx = state_dp.NewContext(state_dp.NewSEMemberState())
	ctx.Request()
	ctx = state_dp.NewContext(state_dp.NewPeMemberState())
	ctx.Request()
}
