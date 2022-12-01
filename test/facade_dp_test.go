package test

import (
	"lbf.com/go-dp/facade_dp"
	"testing"
)

func TestFacadeDp(t *testing.T) {
	facade := facade_dp.BuyFacade{}
	facade.BuyGoods("9527", 10)
}
