package sdktrade

import (
	. "trade/showapiSDK/normalRequest"
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
