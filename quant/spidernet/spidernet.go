package spidernet

import (
	"fmt"
	"trade/config"
	"trade/quant"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/shopspring/decimal"
)

var (
	huobisystem = quant.HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
	huobistock = quant.HuobiStock{
		Client: new(client.MarketClient).Init(config.Host),
	}
	huobialgo = quant.HuobiAlgo{
		Client: new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host),
	}

	// systemInfo quant.SystemInfo
)

func (*SpiderNet) GetSystemStatus(systemInfo quant.SystemInfo) {
	_, err := systemInfo.GetSystemStatus()
	if err != nil {
		panic("Trade System Not Avalable.")
	}
}

func (*SpiderNet) GetAllUsdtTradeSymbols(systemInfo quant.SystemInfo) map[string]map[string]interface{} {
	symbols, err := systemInfo.GetAllUsdtTradeSymbols()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	}
	return symbols
}

func (s *SpiderNet) Update30DAllStockCandleStick(systemInfo quant.SystemInfo, stockInfo quant.StockInfo) {
	s.GetSystemStatus(systemInfo)
	allstocks := s.GetAllUsdtTradeSymbols(systemInfo)
	goodstocks := map[string]decimal.Decimal{}
	for stock, value := range allstocks {
		if value["state"].(string) == "online" {
			goodstocks[stock] = value["min-order-value"].(decimal.Decimal)
		}
	}
	fmt.Println(goodstocks)

}

func Run() {
	ss := SpiderNet{}
	ss.Update30DAllStockCandleStick(&huobisystem, &huobistock)
}
