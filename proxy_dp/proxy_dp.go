package proxy_dp

import (
	"errors"
	"log"
)

type Resource string

type ResourceService interface {
	GetResource(id string) (Resource, error)
}

type ResourceServiceImpl struct {
}

// /220506/010311-16517701917d43.jpg
func (receiver *ResourceServiceImpl) GetResource(id string) (Resource, error) {
	log.Println("执行原方法逻辑")
	return Resource(id), nil
}

type RemoteProxyResourceService struct {
	ResourceService
}

func (r *RemoteProxyResourceService) GetResource(id string) (Resource, error) {
	log.Default().Println("执行代理方法逻辑")
	if id == "" {
		return "", errors.New("参数不能为空")
	}

	return r.ResourceService.GetResource(id)

}

func CreateResourceService() ResourceService {

	return &RemoteProxyResourceService{&ResourceServiceImpl{}}
}
