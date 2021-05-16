package example

import (
	"encoding/json"
	"fmt"
	"trade/config"
	"trade/trade"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

func X() {
	// Market_GetMarketStatus()
	// Market_GetAllSymbols()
	// Martket_GetAllUsdtTradeSymbols()
	Market_GetCurrencys()
}

var (
	huobisystem trade.HuobiSystem = trade.HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
)

func Market_GetMarketStatus() {
	fmt.Println(huobisystem.GetSystemStatus())
}

func Market_GetAllSymbols() {
	aaa := huobisystem.GetAllSymbols()
	aaaj, _ := json.Marshal(aaa)
	fmt.Println(string(aaaj)[0])
}

func Market_GetAllUsdtTradeSymbols() {
	aaa := huobisystem.GetAllUsdtTradeSymbols()
	fmt.Println(aaa)
}

func Market_GetCurrencys() {
	aaa := huobisystem.GetCurrencys()
	fmt.Println(aaa)
}
