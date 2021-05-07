package sdktrade

import (
	"encoding/json"
	"fmt"
	"trade/stockindex/rsi"
)

const (
	ShishiHangqingPiliang = "131-46"
	ShishiHangqing        = "131-44"
	LishiHangqing         = "131-47"
	ShishiFenshixian      = "131-49"
	ShishiKxian           = "131-50"
	Rsi                   = "131-67"
	Appid                 = 621650                             //要替换成自己的
	Sign                  = "ba70e5462ef74adfb9fdd33c70675ba8" //要替换成自己的

	LishiRixian = "1529-2"
)

func TradeTest() {
	// showapi_appid := 621650                            //要替换成自己的
	// showapi_sign := "ba70e5462ef74adfb9fdd33c70675ba8" //要替换成自己的
	// res := ShowapiRequest("http://route.showapi.com/131-44", showapi_appid, showapi_sign)
	// res.AddTextPara("code", "600887")
	// res.AddTextPara("need_k_pic", "0")
	// res.AddTextPara("needIndex", "0")
	// fmt.Println(res.Post())

	// res.AddTextPara("begin", "2016-09-01")
	// res.AddTextPara("end", "2016-09-02")
	// res.AddTextPara("code", "600004")
	// res.AddTextPara("type", "bfq")

}

var (
	auth = NewAuth(Appid, Sign)
)

func TradeTest2() {

	res := auth.SetClient(ShishiHangqing)
	res.AddTextPara("code", "600887")
	res.AddTextPara("need_k_pic", "1")
	res.AddTextPara("needIndex", "0")
	fmt.Println(res.Post())
}

func TradeTest3() {

	res := auth.SetClient(LishiRixian)
	res.AddTextPara("begin", "2016-09-01")
	res.AddTextPara("end", "2016-10-01")
	res.AddTextPara("code", "600004")
	fmt.Println(res.Post())
}

func RsiTest() {
	res := auth.SetClient(Rsi)
	res.AddTextPara("code", "000651")
	res.AddTextPara("start", "20200407")
	res.AddTextPara("end", "20200507")
	res.AddTextPara("real", "1")
	resp, err := res.Post()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp)

	aaa := rsi.RSIResp{}
	err = json.Unmarshal([]byte(resp), &aaa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aaa.ResBody.List[0].Rsi24)
}

// http://blog.codeg.cn/2014/12/27/gochart-info/
