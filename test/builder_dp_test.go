package test

import (
	"lbf.com/go-dp/builder_dp"
	"testing"
)

func TestBuilderDp(t *testing.T) {
	product, err := builder_dp.CreateProductBuilder().
		SetProductTitle("").
		SetProductNum(-1).
		SetProductNormalPrice(2000).
		SetProductGroupPrice(1698).
		SetProductFreightPayer(builder_dp.Seller).
		Build()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(product)
}
