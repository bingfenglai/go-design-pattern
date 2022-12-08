package bridge_dp

type IAuthc interface {
	Authc()
}

type Authenticator struct {
	oauthc OAuthc
}

func NewAuthenticator(OAuthc OAuthc) *Authenticator {
	return &Authenticator{oauthc: OAuthc}
}

func (a *Authenticator) Authc() {
	println("获取认证前的必要信息")
	a.oauthc.OAuthc()
}
