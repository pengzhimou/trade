package trade

import (
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

//不同的交易中心
type MarketInfo interface {
	GetAllSymbols()
}

//Huobi交易中心
type HuobiMarket struct {
	Client *client.CommonClient
}

func (hm *HuobiMarket) GetAllSymbols() []common.Symbol {
	symbols, err := hm.Client.GetSymbols()
	if err != nil {
		klog.Error("Error of GetSymbols!", err)
	}
	return symbols
}

/////////////////////////////////////////////
//不同的交易中心
type StockInfo interface {
	GetLast24hCandlestick()
	GetCandleStick()
	GetBCN()
}

//Huobi股票
type HuobiStock struct {
	Client *client.MarketClient
	Stock  string
}

func (hs *HuobiStock) GetCandleStick() {

}

func (hs *HuobiStock) GetBCN() {

}
