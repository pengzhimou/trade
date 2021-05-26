package spidernet

import (
	"fmt"
	"trade/config"
	. "trade/quant"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/shopspring/decimal"
)

var (
	huobisystem = HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
	huobimarket = HuobiMarket{
		Client: new(client.MarketClient).Init(config.Host),
	}
	huobialgo = HuobiAlgo{
		Client: new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host),
	}

	// systemInfo quant.SystemInfo
)

func (*SpiderNet) GetSystemStatus(systemInfo SystemInfo) {
	_, err := systemInfo.GetSystemStatus()
	if err != nil {
		panic("Trade System Not Avalable.")
	}
}

func (*SpiderNet) GetAllUsdtTradeSymbols(systemInfo SystemInfo) map[string]map[string]interface{} {
	symbols, err := systemInfo.GetAllUsdtTradeSymbols()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	}
	return symbols
}

func (s *SpiderNet) GetAllAvaliableStock(systemInfo SystemInfo, stockInfo MarketInfo) map[string]decimal.Decimal {
	s.GetSystemStatus(systemInfo)
	allstocks := s.GetAllUsdtTradeSymbols(systemInfo)
	goodstocks := map[string]decimal.Decimal{}
	for stock, value := range allstocks {
		if value["state"].(string) == "online" {
			goodstocks[stock] = value["min-order-value"].(decimal.Decimal)
		}
	}
	return goodstocks
}

func (s *SpiderNet) GetStock24HRiseFall(stk Stock, mktInfo MarketInfo) *market.Candlestick {
	candle, err := mktInfo.GetCandleStick24H(stk.Name)
	if err != nil {
		applogger.Error("Get GetStock24HRiseFall error: %s", err)
	}
	return candle

}

func Run() {
	ss := SpiderNet{}
	stocks := ss.GetAllAvaliableStock(&huobisystem, &huobimarket)
	for stock := range stocks {
		stk := Stock{
			Name: stock,
		}
		fmt.Println(ss.GetStock24HRiseFall(stk, &huobimarket))
	}
}
