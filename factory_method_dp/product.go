package factory_method_dp

import "github.com/shopspring/decimal"

type Operation interface {
	GetResult(x, y float64) float64
}

type OperationAdd struct {
}

func (receiver *OperationAdd) GetResult(x, y float64) float64 {
	dx := decimal.NewFromFloat(x)
	dy := decimal.NewFromFloat(y)
	return dx.Add(dy).InexactFloat64()
}

type OperationSub struct {
}

func (r *OperationSub) GetResult(x, y float64) float64 {
	dx := decimal.NewFromFloat(x)
	dy := decimal.NewFromFloat(y)

	return dx.Sub(dy).InexactFloat64()
}

type OperationMul struct {
}

func (o *OperationMul) GetResult(x, y float64) float64 {
	dx := decimal.NewFromFloat(x)
	dy := decimal.NewFromFloat(y)

	return dx.Mul(dy).InexactFloat64()
}

type OperationDiv struct {
}

func (o *OperationDiv) GetResult(x, y float64) float64 {
	dx := decimal.NewFromFloat(x)
	dy := decimal.NewFromFloat(y)

	return dx.Div(dy).InexactFloat64()
}
