package facade_dp

import (
	"fmt"
	"log"
)

type storeService struct {
}

func (receiver *storeService) checkStore(goodsId string, total int) (bool, error) {
	log.Default().Println(fmt.Sprintf("检查%s库存是否充足", goodsId))
	return true, nil
}

type orderService struct {
}

func (receiver orderService) createOrder(goodsId string, total int) (string, error) {
	log.Println("创建订单")

	return goodsId, nil
}

func (receiver orderService) updateOrderState(orderId, state string) {
	log.Println("更新订单状态")
}

func (receiver orderService) getOrder(orderId string) string {
	log.Println("获取订单信息")
	return ""
}

type payService struct {
}

func (receiver *payService) pay(orderId string) (bool, error) {
	log.Println("调用支付网关进行支付" + orderId)
	return true, nil
}

type logisticsService struct {
}

func (receiver logisticsService) logistics(order string) {
	log.Println("创建物流单并发货")
}
