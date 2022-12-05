package singleton_dp

import (
	"log"
	"sync"
)

var (
	instance, instanceV2 *Singleton
	lock                 = &sync.Mutex{}
	once                 = sync.Once{}
)

type Singleton struct {
}

func GetInstance() *Singleton {
	log.Default().Println("获取实例")
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
			log.Default().Println("创建实例")
		}
	}
	return instance
}

func GetInstanceV2() *Singleton {
	log.Default().Println("获取实例")
	//
	once.Do(func() {
		instanceV2 = &Singleton{}
		log.Default().Println("创建实例")
	})
	return instanceV2
}
