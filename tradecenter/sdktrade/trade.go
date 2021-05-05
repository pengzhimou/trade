package sdktrade

import (
	"fmt"
	. "trade/showapiSdk/normalRequest"
)

func TradeTest() {
	showapi_appid := 621650                            //要替换成自己的
	showapi_sign := "ba70e5462ef74adfb9fdd33c70675ba8" //要替换成自己的
	res := ShowapiRequest("http://route.showapi.com/131-44", showapi_appid, showapi_sign)
	res.AddTextPara("code", "600887")
	res.AddTextPara("need_k_pic", "0")
	res.AddTextPara("needIndex", "0")
	fmt.Println(res.Post())

	res.AddTextPara("begin", "2016-09-01")
	res.AddTextPara("end", "2016-09-02")
	res.AddTextPara("code", "600004")
	res.AddTextPara("type", "bfq")

}
