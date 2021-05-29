package spidernet

import (
	"regexp"
	q "trade/quant"
	. "trade/utils"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/shopspring/decimal"
)

func (*SpiderNet) GetSystemStatus(systemInfo q.SystemInfo) {
	_, err := systemInfo.GetSystemStatus()
	if err != nil {
		panic("Trade System Not Avalable.")
	}
}

func (*SpiderNet) GetAllUsdtTradeSymbols(systemInfo q.SystemInfo) map[string]map[string]interface{} {
	symbols, err := systemInfo.GetAllUsdtTradeSymbols()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	}
	return symbols
}

func (s *SpiderNet) GetAllAvaliableStock(systemInfo q.SystemInfo, stockInfo q.MarketInfo) map[string]decimal.Decimal {
	s.GetSystemStatus(systemInfo)
	allstocks := s.GetAllUsdtTradeSymbols(systemInfo)
	goodstocks := map[string]decimal.Decimal{}
	re, _ := regexp.Compile(`\d+[s|l]`)

	for stock, value := range allstocks {
		//黑名单过滤，如果为黑名单内的，sw置off
		sw := "on"
		for _, s := range Blacklist() {
			if stock == s {
				sw = "off"
			}
		}
		//判断非黑名单而来 且 不是加倍的stock
		if sw == "on" && re.FindString(stock) == "" { //搜不出
			if value["state"].(string) == "online" {
				goodstocks[stock] = value["min-order-value"].(decimal.Decimal)
			}
		}
	}
	return goodstocks
}

func getStockRF(stk q.Stock, mktInfo q.MarketInfo, period string) decimal.Decimal {
	candlesRiseFail, err := mktInfo.GetCandleStick(stk.Name, period, 1)
	if err != nil {
		applogger.Error("Get GetStock24HRiseFall error: %s", err)
	}
	riseFail := Cal(Cal(candlesRiseFail[0].Open, candlesRiseFail[0].Close, "-"), candlesRiseFail[0].Open, "/")
	return riseFail
}

func (s *SpiderNet) GetStockRiseFall(stk q.Stock, mktInfo q.MarketInfo, period string) q.Stock {
	riseFailDay := getStockRF(stk, mktInfo, period)
	stk.RiseFailDay = Cal(riseFailDay, decimal.NewFromInt(100), "*")
	return stk
}
