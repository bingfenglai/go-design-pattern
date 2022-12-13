package test

import (
	"lbf.com/go-dp/mediator_dp"
	"testing"
)

func TestMediatorDp(t *testing.T) {
	auctionMediator := mediator_dp.NewLuanxinghaiAuctionMediator()
	hanli := mediator_dp.NewHanli(auctionMediator)
	ziling := mediator_dp.NewZiling(auctionMediator)
	auctionMediator.Immortal2 = ziling
	auctionMediator.Immortal1 = hanli

	hanli.TakeOutWeaponAndDeal()
	ziling.TakeOutWeaponAndDeal()

}
