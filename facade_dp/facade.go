package facade_dp

type BuyFacade struct {
	storeService
	orderService
	payService
	logisticsService
}

func (receiver BuyFacade) BuyGoods(goodsId string, total int) (bool, error) {
	receiver.checkStore(goodsId, total)
	orderId, _ := receiver.createOrder(goodsId, total)
	receiver.pay(orderId)
	receiver.updateOrderState(orderId, "已支付")
	order := receiver.getOrder(orderId)
	receiver.logistics(order)
	return true, nil
}
