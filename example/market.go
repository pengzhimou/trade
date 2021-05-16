package example

import (
	"encoding/json"
	"fmt"
	"trade/config"
	"trade/trade"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

func Y() {
	// Market_GetCandleStick()
	Market_GetCandleStick24H()
}

var (
	huobistock = trade.HuobiStock{
		Client: new(client.MarketClient).Init(config.Host),
	}
	stock = "oneusdt"
	min5  = market.MIN5
	min15 = market.MIN15
	day1  = market.DAY1
)

func P(s interface{}) string {
	aaa, _ := json.Marshal(s)
	rst := string(aaa)
	fmt.Println(rst)
	return (rst)
}

func Market_GetCandleStick() {
	a5, _ := huobistock.GetCandleStick(stock, min5, 5)
	P(a5)
}
func Market_GetCandleStick24H() {
	a5, _ := huobistock.GetCandleStick24H(stock)
	P(a5)
}
