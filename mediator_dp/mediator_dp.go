package mediator_dp

import "log"

// 修仙者 抽象的同事角色
type Immortal interface {
	// 拿出宝物交换
	TakeOutWeaponAndDeal()
	// 获取交换得到的宝物
	GainWeapon(weapon string)
}

// 韩立
type Hanli struct {
	mediator AuctionMediator
}

func NewHanli(mediator AuctionMediator) *Hanli {
	return &Hanli{mediator: mediator}
}

func (receiver *Hanli) TakeOutWeaponAndDeal() {
	log.Default().Println("韩立拿出了千年雪灵芝")
	receiver.mediator.Business("千年雪灵芝", receiver)
}

func (receiver *Hanli) GainWeapon(weapon string) {
	log.Default().Printf("韩立得到%s", weapon)
}

// 紫菱仙子
type Ziling struct {
	mediator AuctionMediator
}

func (receiver *Ziling) TakeOutWeaponAndDeal() {
	log.Default().Println("紫菱仙子拿出了金雷竹")
	receiver.mediator.Business("金雷竹", receiver)
}

func (receiver *Ziling) GainWeapon(weapon string) {
	log.Default().Printf("紫菱仙子得到%s", weapon)
}

func NewZiling(mediator AuctionMediator) *Ziling {
	return &Ziling{mediator: mediator}
}

// 法宝拍卖会中介者角色
type AuctionMediator interface {

	// 修仙者交易法宝方法
	Business(weapon string, immortal Immortal)
}

// 具体拍卖会
type LuanxinghaiAuctionMediator struct {
	Immortal1 Immortal
	Immortal2 Immortal
}

func NewLuanxinghaiAuctionMediator() *LuanxinghaiAuctionMediator {
	return &LuanxinghaiAuctionMediator{}
}

func (receiver *LuanxinghaiAuctionMediator) Business(weapon string, immortal Immortal) {
	if immortal == receiver.Immortal1 {
		receiver.Immortal2.GainWeapon(weapon)
	}

	if immortal == receiver.Immortal2 {
		receiver.Immortal1.GainWeapon(weapon)
	}
}
