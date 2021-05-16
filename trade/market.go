package trade

import "github.com/huobirdcenter/huobi_golang/pkg/client"

//不同的交易中心
type MarketInfo interface {
	GetAllSymbols()
}

//Huobi交易中心
type HuobiMarket struct {
	Client *client.CommonClient
}

func (hm *HuobiMarket) GetAllSymbols() {
	hm.Client.GetSymbols()

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
