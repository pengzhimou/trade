package quant

import (
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

/////////////////////////////////////////////
//不同的交易中心
type MarketInfo interface {
	GetCandleStick24H(stock string) (*market.Candlestick, error)
	GetCandleStick(stock string, period string, size int) ([]market.Candlestick, error)
	GetBuySellTick(stock string, count int) (*market.Depth, error)
	GetLatestTrade(stock string) (*market.TradeTick, error)
}

//Huobi股票
type HuobiMarket struct {
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
// [{
// "amount": "152888.14588048446",
// "close": "0.149563",
// "count": 63,
// "high": "0.149563",
// "id": 1621155300,
// "low": "0.1488",
// "open": "0.149",
// "vol": "22810.71191588"
// }]
//	client := new(client.MarketClient).Init(config.Host)
func (hs *HuobiMarket) GetCandleStick(stock string, period string, size int) ([]market.Candlestick, error) {
	optionalRequest := market.GetCandlestickOptionalRequest{Period: period, Size: size}
	resp, err := hs.Client.GetCandlestick(stock, optionalRequest)
	if err != nil {
		applogger.Error("GetCandlestick error", err.Error())
	}
	return resp, err
}

//
// 与上面一致
func (hs *HuobiMarket) GetCandleStick24H(stock string) (*market.Candlestick, error) {
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
// 	Bids      [][]decimal.Decimal `json:"bids"`   //买 高-低
// 	Asks      [][]decimal.Decimal `json:"asks"`   //卖 低-高
// }
//Tips: count最小是5 最大20
// {
// 	"asks": [
// 	  [
// 		"0.159545",
// 		"3760.69"
// 	  ],
// 	  [
// 		"0.159546",
// 		"1768.48"
// 	  ],
// 	],
// 	"bids": [
// 	  [
// 		"0.159308",
// 		"1900"
// 	  ],
// 	  [
// 		"0.159283",
// 		"3766.88"
// 	  ],
// 	],
// 	"ts": "1621171160001",
// 	"version": "101027481200"
//   }
func (hs *HuobiMarket) GetBuySellTick(stock string, count int) (*market.Depth, error) {
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
// {
// 	"Data": [
// 	  {
// 		"amount": "51.41",
// 		"direction": "sell",
// 		"id": "101027500699277477829795329",
// 		"price": "0.15937",
// 		"trade-id": 6425270,
// 		"ts": "1621171295037"
// 	  }
// 	],
// 	"id": "101027500699",
// 	"ts": "1621171295037"
//   }]
func (hs *HuobiMarket) GetLatestTrade(stock string) (*market.TradeTick, error) {
	resp, err := hs.Client.GetLatestTrade(stock)
	return resp, err
}
