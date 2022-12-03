package builder_dp

import (
	"time"
)

type FreightPayer string

const (
	Seller = "Seller"
	Buyer  = "buyer"
)

type Product struct {
	Title string
	Price
	CatId int
	Desc  string
	Num   int
	// 邮费
	DeliveryFee
	// 运费承担方式 seller（卖家承担）,buyer(买家承担)
	FreightPayer
	// 上架时间
	ListTime *time.Time
	// 支持会员打折
	HasDiscount bool
	// 商品的积分返点比例。如:5,表示:返点比例0.5%.
	//注意：返点比例必须是>0的整数，而且最大是90,即为9%.
	AuctionPoint int
}

type Price struct {
	NormalPrice int
	GroupPPrice int
}

type DeliveryFee struct {
	// 平邮费用
	PostFee int
	// 快递费用
	ExpressFee int
	EmsFee     int
}

type ProductBuilder struct {
	*Product
	*buildErrorWrapper
}

func CreateProductBuilder() *ProductBuilder {
	return &ProductBuilder{Product: &Product{}}
}

func (receiver *ProductBuilder) newError(msg string) {
	if receiver.buildErrorWrapper == nil {
		receiver.buildErrorWrapper = &buildErrorWrapper{}
	}
	receiver.buildErrorWrapper.newError(msg)
}

func (receiver *ProductBuilder) SetProductTitle(title string) *ProductBuilder {
	if title == "" || len(title) > 64 {
		receiver.newError("产品标题不能为空或大于64个字符")
	}
	receiver.Product.Title = title
	return receiver
}

func (receiver *ProductBuilder) SetProductNormalPrice(normalPrice int) *ProductBuilder {
	receiver.Product.NormalPrice = normalPrice
	return receiver
}

func (receiver *ProductBuilder) SetProductGroupPrice(groupPrice int) *ProductBuilder {
	receiver.Product.GroupPPrice = groupPrice
	return receiver
}

func (receiver *ProductBuilder) CatId(catId int) *ProductBuilder {
	receiver.Product.CatId = catId
	return receiver
}

func (receiver *ProductBuilder) Desc(desc string) *ProductBuilder {
	receiver.Product.Desc = desc
	return receiver
}

func (receiver *ProductBuilder) SetProductNum(num int) *ProductBuilder {

	if num < 0 {
		receiver.newError("产品数量不能小于0")
	}

	receiver.Product.Num = num
	return receiver
}

func (receiver *ProductBuilder) PostFee(postFee int) *ProductBuilder {
	receiver.Product.PostFee = postFee
	return receiver
}

func (receiver *ProductBuilder) ExpressFee(expressFree int) *ProductBuilder {
	receiver.Product.ExpressFee = expressFree
	return receiver
}

func (receiver *ProductBuilder) EmsFee(emsFee int) *ProductBuilder {
	receiver.Product.EmsFee = emsFee
	return receiver
}

func (receiver *ProductBuilder) SetProductFreightPayer(freightPayer FreightPayer) *ProductBuilder {
	receiver.Product.FreightPayer = freightPayer
	return receiver
}

func (receiver *ProductBuilder) ListTime(listTime *time.Time) *ProductBuilder {
	receiver.Product.ListTime = listTime
	return receiver
}

func (receiver *ProductBuilder) HasDiscount(hasDiscount bool) *ProductBuilder {
	receiver.Product.HasDiscount = hasDiscount
	return receiver
}

func (receiver *ProductBuilder) AuctionPoint(auctionPoint int) *ProductBuilder {
	receiver.Product.AuctionPoint = auctionPoint
	return receiver
}

func (receiver ProductBuilder) Build() (*Product, error) {
	if receiver.error != nil {
		return nil, receiver.error
	}

	return receiver.Product, nil
}
