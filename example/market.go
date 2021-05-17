package example

import (
	"fmt"
	"trade/config"
	"trade/quant"
	"trade/utils"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

func Y() {
	// Market_GetCandleStick()
	// Market_GetCandleStick24H()
	// Market_GetBuySellTick()
	// Market_GetLatestTrade()
	Market_Bind()
}

var (
	huobistock = quant.HuobiStock{
		Client: new(client.MarketClient).Init(config.Host),
	}
	stock = "oneusdt"
	min5  = market.MIN5
	min15 = market.MIN15
	day1  = market.DAY1
)

func Market_GetCandleStick() {
	a5, _ := huobistock.GetCandleStick(stock, min5, 5)
	utils.P(a5)
}

func Market_GetCandleStick24H() {
	a5, _ := huobistock.GetCandleStick24H(stock)
	utils.P(a5)
}

func Market_GetBuySellTick() {
	a5, _ := huobistock.GetBuySellTick(stock, 5)
	utils.P(a5)
}

func Market_GetLatestTrade() {
	a5, _ := huobistock.GetLatestTrade(stock)
	utils.P(a5)
}

func Market_Bind() {
	a5, _ := huobistock.GetBuySellTick(stock, 5)
	fmt.Println(a5.Asks[0], a5.Bids[0])
	a6, _ := huobistock.GetLatestTrade(stock)
	fmt.Println(a6.Data)
}
