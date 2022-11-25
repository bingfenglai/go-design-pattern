// 第一版策略模式demo,不涉及其他设计模式混合使用
package strategy_dp_v1

// 收费策略
type FeesStrategy interface {
	Settlement(total float64) float64
}

// 计费上下文
type FeesContext struct {
	FeesStrategy
}

func (receiver FeesContext) Settlement(total float64) float64 {
	return receiver.FeesStrategy.Settlement(total)
}

func CreateFreesContext(fst FeesStrategy) *FeesContext {

	return &FeesContext{fst}
}
