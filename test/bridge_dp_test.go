package test

import (
	"lbf.com/go-dp/bridge_dp"
	"testing"
)

func TestBridgeDp(t *testing.T) {
	authenticator := bridge_dp.NewAuthenticator(&bridge_dp.WechatAuthc{})
	authenticator.Authc()
	authenticator = bridge_dp.NewAuthenticator(&bridge_dp.QqAuthc{})
	authenticator.Authc()
}
