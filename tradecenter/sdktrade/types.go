package sdktrade

import (
	. "trade/showapiSdk/normalRequest"
)

const (
	ShishiHangqingPiliang = "131-46"
	ShishiHangqing        = "131-44"
	LishiHangqing         = "131-47"
	ShishiFenshixian      = "131-49"
	ShishiKxian           = "131-50"
)

type Auth struct {
	showapi_appid int
	showapi_sign  string
}

func NewAuth(appid int, token string) *Auth {
	return &Auth{
		showapi_appid: appid,
		showapi_sign:  token,
	}
}

func (a *Auth) SetClient(uri string) *NormalReq {
	return ShowapiRequest("http://route.showapi.com/"+uri, a.showapi_appid, a.showapi_sign)
}
