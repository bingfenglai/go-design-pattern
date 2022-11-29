package factory_dp

import "errors"

type OptionFactory interface {
	CreateOption(name string) (Option, error)
}

type OptionFactoryImpl struct {
}

func GetFactoryInstance() OptionFactory {
	return &OptionFactoryImpl{}
}

func (receiver *OptionFactoryImpl) CreateOption(name string) (Option, error) {

	switch name {
	case "+":
		return &OptionAdd{}, nil
	case "-":
		return &OptionSub{}, nil
	case "*":
		return &OptionMul{}, nil
	case "/":
		return &OptionDiv{}, nil
	default:
		return nil, errors.New("输入参数不正确")

	}
}
