package trade

import (
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

func RunAllExamplesSystem() {
	// getSystemStatus()
	// getMarketStatus()

	// getSymbols()
	//// 2021-05-09T17:09:01.787+0800	INFO	symbol: qashbtc, {BaseCurrency:qash QuoteCurrency:btc PricePrecision:8 AmountPrecision:4 SymbolPartition:innovation Symbol:qashbtc State:online ValuePrecision:8 LimitOrderMinOrderAmt:0.1 LimitOrderMaxOrderAmt:17000000 SellMarketMinOrderAmt:0.1 SellMarketMaxOrderAmt:1700000 BuyMarketMaxOrderValue:9 MinOrderValue:0.0001 MaxOrderValue:0 LeverageRatio:0}

	// getCurrencys()
	// // 2021-05-09T17:10:42.275+0800	INFO	Get currency, count=404
	// // 2021-05-09T17:10:42.275+0800	INFO	currency: usdt
	// // 2021-05-09T17:10:42.275+0800	INFO	currency: btc
	// // 2021-05-09T17:10:42.276+0800	INFO	currency: bch
	// // 2021-05-09T17:10:42.276+0800	INFO	currency: eth

	// getV2ReferenceCurrencies()
	// getTimestamp()
}

func getSystemStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSystemStatus()
	if err != nil {
		applogger.Error("Get system status error: %s", err)
	} else {
		applogger.Info("Get system status %s", resp)
	}
}

func getMarketStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetMarketStatus()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	} else {
		applogger.Info("Get market status, status: %d", resp.MarketStatus)
	}
}

func getSymbols() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSymbols()
	if err != nil {
		applogger.Error("Get symbols error: %s", err)
	} else {
		applogger.Info("Get symbols, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("symbol: %s, %+v", result.Symbol, result)
		}
	}
}

func getCurrencys() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetCurrencys()

	if err != nil {
		applogger.Error("Get currency error: %s", err)
	} else {
		applogger.Info("Get currency, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("currency: %+v", result)
		}
	}
}

func getV2ReferenceCurrencies() {
	optionalRequest := common.GetV2ReferenceCurrencies{Currency: "", AuthorizedUser: "true"}

	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetV2ReferenceCurrencies(optionalRequest)

	if err != nil {
		applogger.Error("Get reference currency error: %s", err)
	} else {
		applogger.Info("Get reference currency, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("currency:%s, ", result.Currency)

			for _, chain := range result.Chains {
				applogger.Info("Chain: %+v", chain)
			}
		}
	}
}

func getTimestamp() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetTimestamp()

	if err != nil {
		applogger.Error("Get timestamp error: %s", err)
	} else {
		applogger.Info("Get timestamp: %d", resp)
	}
}
