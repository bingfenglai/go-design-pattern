package decorator_dp

import "log"

// api接口
type Api interface {

	// 访问Api的方法
	Access()
}

type GetResourceApi struct {
}

func (receiver *GetResourceApi) Access() {

	log.Println("访问api接口")
}

type Decorator interface {
	Api
}

type SecurityDecorator struct {
	Decorator
}

func (receiver *SecurityDecorator) Access() {
	log.Default().Println("安全的api")
	receiver.Decorator.Access()
}

type EncryptionDecorator struct {
	Decorator
}

func (receiver *EncryptionDecorator) Access() {
	log.Default().Println("加密的api")
	receiver.Decorator.Access()
}
