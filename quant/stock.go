package quant

import "github.com/huobirdcenter/huobi_golang/pkg/model/market"

type BuySell struct {
}

type Stock struct {
	Name         string
	CandleDay    []market.Candlestick
	CandleVMin   []market.Candlestick
	CandleVVVMin []market.Candlestick
	BuySellTick  []BuySell
}

func (s *Stock) DrawVMin(count int) {

}

func (s *Stock) DrawVVVMin(count int) {

}

func (s *Stock) DrawDay(count int) {

}

func (s *Stock) DrawBuySellBalance() {

}
