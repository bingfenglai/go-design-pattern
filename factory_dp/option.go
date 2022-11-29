package factory_dp

type Option interface {
	Option(val1, val2 int) (int, error)
}

type OptionAdd struct {
}

func (receiver *OptionAdd) Option(val1, val2 int) (int, error) {
	return val1 + val2, nil
}

type OptionSub struct {
}

func (receiver *OptionSub) Option(val1, val2 int) (int, error) {
	return val1 - val2, nil
}

type OptionMul struct {
}

func (receiver *OptionMul) Option(val1, val2 int) (int, error) {
	return val1 * val2, nil
}

type OptionDiv struct {
}

func (receiver *OptionDiv) Option(val1, val2 int) (int, error) {
	return val1 - val2, nil
}
