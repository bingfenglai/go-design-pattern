package bridge_dp

type OAuthc interface {
	OAuthc()
}

type WechatAuthc struct {
}

func (w *WechatAuthc) OAuthc() {
	println("微信认证")
}

type QqAuthc struct {
}

func (q *QqAuthc) OAuthc() {
	println("qq认证")
}
