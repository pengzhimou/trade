package trade

import (
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

//不同的交易中心
type SystemInfo interface {
	GetAllSymbols() []common.Symbol
	GetSystemStatus() *common.MarketStatus
	GetAllUsdtTradeSymbols() map[string]map[string]interface{}
}

//Huobi交易中心
type HuobiSystem struct {
	Client *client.CommonClient
}

// [{
//     "amount-precision": 4,
//     "base-currency": "nhbtc",
//     "buy-market-max-order-value": "100000",
//     "leverage-ratio": "0",
//     "limit-order-max-order-amt": "1111111",
//     "limit-order-min-order-amt": "0.01",
//     "max-order-value": "0",
//     "min-order-value": "5",
//     "price-precision": 4,
//     "quote-currency": "usdt",
//     "sell-market-max-order-amt": "111111",
//     "sell-market-min-order-amt": "0.01",
//     "state": "online",
//     "symbol": "nhbtcusdt",
//     "symbol-partition": "potentials",
//     "value-precision": 8
//   },]
func (hs *HuobiSystem) GetAllSymbols() ([]common.Symbol, error) {
	symbols, err := hs.Client.GetSymbols()
	if err != nil {
		applogger.Error("Error of GetSymbols!", err)
	}
	return symbols, err
}

func (hs *HuobiSystem) GetAllUsdtTradeSymbols() (map[string]map[string]interface{}, error) {
	allSymbols, err := hs.GetAllSymbols()
	if err != nil {
		return nil, nil
	}
	usdtSymbls := map[string]map[string]interface{}{}
	for _, symbol := range allSymbols {
		if symbol.QuoteCurrency == "usdt" {
			usdtSymbls[symbol.Symbol] = map[string]interface{}{
				"max-order-value":            symbol.MaxOrderValue,
				"min-order-value":            symbol.MinOrderValue,
				"buy-market-max-order-value": symbol.BuyMarketMaxOrderValue,
				"state":                      symbol.State,
			}
		}
	}
	return usdtSymbls, err
}

func (hs *HuobiSystem) GetSystemStatus() (*common.MarketStatus, error) {
	resp, err := hs.Client.GetMarketStatus()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	} else {
		applogger.Info("Get market status, status: %d", resp.MarketStatus)
	}
	return resp, err
}

func (hs *HuobiSystem) GetCurrencys() ([]string, error) {
	resp, err := hs.Client.GetCurrencys()
	if err != nil {
		applogger.Error("Get currency error: %s", err)
	}
	return resp, err
}
