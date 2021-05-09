package trade

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

func RunAllExamplesMarket() {
	// getCandlestick("cnnsusdt", market.MIN5, 5)
	getLast24hCandlestickAskBid()
	getLast24hCandlesticks()
	// getDepth()
	// getLatestTrade()
	// getHistoricalTrade()
	// getLast24hCandlestick()
}

//  Get the candlestick/kline for the stock.
//  period: market.DAY1 market.MIN1 market.MIN5 market.MIN15 and so on
func getCandlestick(stock, period string, size int) {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: period, Size: size}

	resp, err := client.GetCandlestick(stock, optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, kline := range resp {
			applogger.Info("Open=%v Close(Now)=%v High=%v Low=%v Vol=%v", kline.Open, kline.Close, kline.High, kline.Low, kline.Vol)
		}
	}
}

//  Get the latest ticker with some important 24h aggregated market data for btcusdt.
func getLast24hCandlestickAskBid() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid("cnnsusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Bid=%+v, Ask=%+v", resp.Bid, resp.Ask)
	}
}

//  Get the latest tickers for all supported pairs
func getLast24hCandlesticks() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, tick := range resp {
			applogger.Info("Symbol: %s, High: %v, Low: %v, Ask[%v, %v], Bid[%v, %v]",
				tick.Symbol, tick.High, tick.Low, tick.Ask, tick.AskSize, tick.Bid, tick.BidSize)
		}
	}
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

//  Get the latest trade with btucsdt price, volume, and direction.
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
