package observer_dp

type Subject interface {
	// 通知方法
	Notify()
	// 增加观察者
	Attach(observer Observer)
	// 移除观察者
	Detach(observer Observer)
}

type UserRegisteredSubject struct {
	observers []Observer
}

func NewUserRegisteredSubject(observers []Observer) *UserRegisteredSubject {
	return &UserRegisteredSubject{observers: observers}
}

func (receiver *UserRegisteredSubject) Notify() {
	for _, observer := range receiver.observers {
		observer.Update()
	}
}

func (receiver *UserRegisteredSubject) Attach(observer Observer) {
	receiver.observers = append(receiver.observers, observer)
}

func (receiver *UserRegisteredSubject) Detach(observer Observer) {
	removeIdx := -1
	for i, obs := range receiver.observers {
		if obs == observer {
			removeIdx = i
			break
		}
	}

	if removeIdx == -1 {
		return
	}

	receiver.observers = append(receiver.observers[:removeIdx], receiver.observers[removeIdx+1:]...)

}
