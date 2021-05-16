package trade

import (
	"fmt"

	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

func RunAllExamplesMarket() {
	client := new(client.MarketClient).Init(config.Host)

	aaa, _ := GetCandlestick(client, "cnnsusdt", market.MIN5, 2)
	// aaa, _ := getLast24hCandlestickAskBid("cnnsusdt")
	// aaa, _ := getLast24hCandlesticks()
	// aaa, _ := getDepth("btcusdt", 5)
	// aaa, _ := getLatestTrade("cnnsusdt")
	// getHistoricalTrade()
	getLast24hCandlestick()

	fmt.Printf("%s\n", aaa)

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
// [
// {119210.8 0.006252 0.00626 0.006276 %!s(int64=1621128000) %!s(int64=11) 0.006252 746.18579133}
// {135996.48 0.006226 0.006245 0.00628 %!s(int64=1621127700) %!s(int64=37) 0.006221 850.93955969}
// ]
//	client := new(client.MarketClient).Init(config.Host)
func GetCandlestick(client *client.MarketClient, stock, period string, size int) ([]market.Candlestick, error) {
	optionalRequest := market.GetCandlestickOptionalRequest{Period: period, Size: size}
	resp, err := client.GetCandlestick(stock, optionalRequest)
	return resp, err
}

// 个股买1卖1
//  Get the latest ticker with some important 24h aggregated market data for btcusdt.
// type CandlestickAskBid struct {
// 	Amount  decimal.Decimal   `json:"amount"`
// 	Open    decimal.Decimal   `json:"open"`
// 	Close   decimal.Decimal   `json:"close"`
// 	High    decimal.Decimal   `json:"high"`
// 	Id      int64             `json:"id"`
// 	Count   int64             `json:"count"`
// 	Low     decimal.Decimal   `json:"low"`
// 	Vol     decimal.Decimal   `json:"vol"`
// 	Version int64             `json:"version"`
// 	Bid     []decimal.Decimal `json:"bid"`
// 	Ask     []decimal.Decimal `json:"ask"`
// }
//&{460691172.2855285 0.008927 0.007344 0.008927 201057703709 29868 0.007111 3619883.082731892 201057703709 [0.007322 3681.01] [0.007343 5152.78]}
func getLast24hCandlestickAskBid(client *client.MarketClient, stock string) (*market.CandlestickAskBid, error) {
	resp, err := client.GetLast24hCandlestickAskBid(stock)
	return resp, err
}

// 全部股票价格，买1卖1感觉没啥用
func getLast24hCandlesticks() ([]market.SymbolCandlestick, error) {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()
	return resp, err
}

// 买N卖N
// type Depth struct {
// 	Timestamp int64               `json:"ts"`
// 	Version   int64               `json:"version"`
// 	Bids      [][]decimal.Decimal `json:"bids"`
// 	Asks      [][]decimal.Decimal `json:"asks"`
// }
// &{1620834698905 127610781550 [[55549.13 0.457766] [55547.79 0.002] [55547.77 0.050946] [55547.24 1.05868] [55546.35 0.050634]] [[55549.14 0.17] [55551.89 0.0435] [55551.9 0.000386] [55551.91 0.0495] [55554.34 0.051]]}
func getDepth(stock string, count int) (*market.Depth, error) {
	optionalRequest := market.GetDepthOptionalRequest{count}
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
func getLatestTrade(stock string) (*market.TradeTick, error) {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLatestTrade(stock)
	return resp, err
}

//  Get the most recent trades with btcusdt price, volume, and direction.
func getHistoricalTrade() {
	client := new(client.MarketClient).Init(config.Host)
	optionalRequest := market.GetHistoricalTradeOptionalRequest{5}
	resp, err := client.GetHistoricalTrade("btcusdt", optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, tradeData := range resp {
			for _, trade := range tradeData.Data {
				applogger.Info("price: %v", trade.Price)
			}
		}
	}
}

//  Get the summary of trading in the market for the last 24 hours.
func getLast24hCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestick("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Close=%v, Open=%v", resp.Close, resp.Open)
	}
}
