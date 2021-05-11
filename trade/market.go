package trade

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

func RunAllExamplesMarket() {
	// getCandlestick("cnnsusdt", market.MIN5, 5)
	// aaa, _ := getLast24hCandlestickAskBid("cnnsusdt")
	// aaa, _ := getLast24hCandlesticks()
	getDepth()
	// getLatestTrade()
	// getHistoricalTrade()
	// getLast24hCandlestick()

	// fmt.Printf("%s\n", aaa[0])

}

// 个股蜡烛图
//  period: market.DAY1 market.MIN1 market.MIN5 market.MIN15 and so on

func GetCandlestick(stock, period string, size int) ([]market.Candlestick, error) {
	client := new(client.MarketClient).Init(config.Host)
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
func getLast24hCandlestickAskBid(stock string) (*market.CandlestickAskBid, error) {
	client := new(client.MarketClient).Init(config.Host)
	resp, err := client.GetLast24hCandlestickAskBid(stock)
	return resp, err
}

// 全部股票价格，买1卖1感觉没啥用
func getLast24hCandlesticks() ([]market.SymbolCandlestick, error) {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()
	return resp, err
}

//  Get the current order book of the btcusdt.
func getDepth() {
	optionalRequest := market.GetDepthOptionalRequest{10}
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetDepth("btcusdt", market.STEP0, optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, ask := range resp.Asks {
			applogger.Info("ask: %+v", ask)
		}
		for _, bid := range resp.Bids {
			applogger.Info("bid: %+v", bid)
		}

	}
}

// 买10卖10
func getLatestTrade() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLatestTrade("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, trade := range resp.Data {
			applogger.Info("Id=%v, Price=%v", trade.Id, trade.Price)
		}
	}
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
