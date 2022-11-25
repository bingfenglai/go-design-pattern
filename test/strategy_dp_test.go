package test

import (
	"lbf.com/go-dp/strategy_dp_v1"
	"testing"
)

func TestStrategyDp(t *testing.T) {
	frees := 8848.98
	t.Log("原价：$", frees)

	freesContext := strategy_dp_v1.CreateFreesContext(&strategy_dp_v1.FreesStrategyNormal{})
	total := freesContext.Settlement(frees)
	t.Log("执行正常计费应收: $", total)

	freesContext.FeesStrategy = &strategy_dp_v1.FreesStrategyDiscount{Discount: 0.80}
	total = freesContext.Settlement(frees)
	t.Log("执行折扣计费应收：$", total)

	freesContext.FeesStrategy = &strategy_dp_v1.FreesStrategyFullDecrement{Full: 5000, Decrement: 300}
	total = freesContext.Settlement(frees)
	t.Log("执行满减计费应收：$", total)

}
