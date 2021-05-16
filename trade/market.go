package trade

import (
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

/////////////////////////////////////////////
//不同的交易中心
type StockInfo interface {
	GetLast24hCandlestick()
	GetCandleStick()
	GetBuyCellTick()
}

//Huobi股票
type HuobiStock struct {
	Client *client.MarketClient
}

// 个股蜡烛图
//  period: market.DAY1 market.MIN1 market.MIN5 market.MIN15 and so on
// type Candlestick struct {
// 	Amount decimal.Decimal `json:"amount"`
// 	Open   decimal.Decimal `json:"open"`
// 	Close  decimal.Decimal `json:"close"`
// 	High   decimal.Decimal `json:"high"`
// 	Id     int64           `json:"id"`
// 	Count  int64           `json:"count"`
// 	Low    decimal.Decimal `json:"low"`
// 	Vol    decimal.Decimal `json:"vol"`
// }
// {
// "amount": "152888.14588048446",
// "close": "0.149563",
// "count": 63,
// "high": "0.149563",
// "id": 1621155300,
// "low": "0.1488",
// "open": "0.149",
// "vol": "22810.71191588"
// }
//	client := new(client.MarketClient).Init(config.Host)
func (hs *HuobiStock) GetCandleStick(stock string, period string, size int) ([]market.Candlestick, error) {
	optionalRequest := market.GetCandlestickOptionalRequest{Period: period, Size: size}
	resp, err := hs.Client.GetCandlestick(stock, optionalRequest)
	if err != nil {
		applogger.Error("GetCandlestick error", err.Error())
	}
	return resp, err
}

//
func (hs *HuobiStock) GetCandleStick24H(stock string) (*market.Candlestick, error) {
	resp, err := hs.Client.GetLast24hCandlestick(stock)
	if err != nil {
		applogger.Error(err.Error())
	}
	return resp, err
}

// 买N卖N
// type Depth struct {
// 	Timestamp int64               `json:"ts"`
// 	Version   int64               `json:"version"`
// 	Bids      [][]decimal.Decimal `json:"bids"`
// 	Asks      [][]decimal.Decimal `json:"asks"`
// }
func (hs *HuobiStock) GetBuyCellTick(stock string, count int) (*market.Depth, error) {
	optionalRequest := market.GetDepthOptionalRequest{Size: count}
	client := new(client.MarketClient).Init(config.Host)
	resp, err := client.GetDepth(stock, market.STEP0, optionalRequest)
	return resp, err
}

// 最新成交价
// type TradeTick struct {
// 	Id   int64 `json:"id"`
// 	Ts   int64 `json:"ts"`
// 	Data []struct {
// 		Amount    decimal.Decimal `json:"amount"`
// 		TradeId   int64           `json:"trade-id"`
// 		Ts        int64           `json:"ts"`
// 		Id        decimal.Decimal `json:"id"`
// 		Price     decimal.Decimal `json:"price"`
// 		Direction string          `json:"direction"`
// 	}
// }
// &{%!s(int64=100533462842) %!s(int64=1620835255045) [{1972.4 %!s(int64=100010984772) %!s(int64=1620835255045) 100533462842276050712135891 0.00771 buy} {5925.27 %!s(int64=100010984771) %!s(int64=1620835255045) 100533462842276050426753811 0.007695 buy}]}
// &{%!s(int64=100533466966) %!s(int64=1620835318700) [{52.43 %!s(int64=100010984871) %!s(int64=1620835318700) 100533466966276046810652863 0.007631 sell} {11984.16 %!s(int64=100010984870) %!s(int64=1620835318700) 100533466966276053363660985 0.007634 sell} {7144.5 %!s(int64=100010984869) %!s(int64=1620835318700) 100533466966276046810651535 0.007635 sell}]}
func (hs *HuobiStock) GetLatestTrade(stock string) (*market.TradeTick, error) {
	resp, err := hs.Client.GetLatestTrade(stock)
	return resp, err
}
