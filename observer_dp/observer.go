package observer_dp

import "log"

type Observer interface {
	Update()
}

type RightObserver struct {
}

func (c *RightObserver) Update() {
	log.Default().Println("接收到订阅的通知,分配初始权益")
}

type IntegralObserver struct {
}

func (receiver *IntegralObserver) Update() {
	log.Default().Println("接收到订阅的通知,初始化积分")
}

type GradeObserver struct {
}

func (receiver GradeObserver) Update() {
	log.Default().Println("接收到订阅的通知,初始化用户等级")

}
