package example

import (
	"trade/config"
	"trade/quant"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

func X() {
	// Market_GetMarketStatus()
	// Market_GetAllSymbols()
	// Martket_GetAllUsdtTradeSymbols()
	Market_GetCurrencys()
}

var (
	huobisystem = quant.HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
)

func Market_GetMarketStatus() {
	a5, _ := huobisystem.GetSystemStatus()
	P(a5)
}

func Market_GetAllSymbols() {
	a5, _ := huobisystem.GetAllSymbols()
	P(a5)
}

func Market_GetAllUsdtTradeSymbols() {
	a5, _ := huobisystem.GetAllUsdtTradeSymbols()
	P(a5)
}

func Market_GetCurrencys() {
	a5, _ := huobisystem.GetCurrencys()
	P(a5)
}
