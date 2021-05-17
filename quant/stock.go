package quant

import "github.com/huobirdcenter/huobi_golang/pkg/model/market"

type BuyCell struct {
}

type Stock struct {
	Name         string
	CandleDay    []market.Candlestick
	CandleVMin   []market.Candlestick
	CandleVVVMin []market.Candlestick
	BuyCellTick  []BuyCell
}

func (s *Stock) DrawVMin(count int) {

}

func (s *Stock) DrawVVVMin(count int) {

}

func (s *Stock) DrawDay(count int) {

}

func (s *Stock) DrawBuyCellBalance() {

}
