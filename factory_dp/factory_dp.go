package factory_dp

import "errors"

type OptionFactory struct {
}

func GetFactoryInstance() *OptionFactory {
	return &OptionFactory{}
}

func (receiver *OptionFactory) CreateOption(name string) (Option, error) {

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
