package strategy_dp_v1

import "github.com/shopspring/decimal"

// 原价计费
type FreesStrategyNormal struct {
}

func (receiver *FreesStrategyNormal) Settlement(total float64) float64 {
	return total
}

// 折扣计费
type FreesStrategyDiscount struct {
	Discount float64
}

func (receiver *FreesStrategyDiscount) Settlement(total float64) float64 {

	f, _ := decimal.NewFromFloat(total).Mul(decimal.NewFromFloat(receiver.Discount)).Round(2).Float64()

	return f
}

// 满减计费
type FreesStrategyFullDecrement struct {
	Full      float64
	Decrement float64
}

func (receiver *FreesStrategyFullDecrement) Settlement(total float64) float64 {

	if receiver.Full <= total {
		return total - receiver.Decrement
	}

	return total

}
