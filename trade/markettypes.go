package trade

import "github.com/huobirdcenter/huobi_golang/pkg/client"

//不同的交易中心
type StockMarket interface {
	GetCandleStick()
	GetBCN()
}
type HuobiMarket struct {
	Client *client.MarketClient
}
