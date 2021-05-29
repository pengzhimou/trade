package quant

import (
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/shopspring/decimal"
)

type BuySell struct {
}

type Stock struct {
	Name           string
	RiseFailDay    decimal.Decimal
	RiseFailVVVMin decimal.Decimal
	RiseFailVMin   decimal.Decimal
	CandleDay      []market.Candlestick
	CandleVVVMin   []market.Candlestick
	CandleVMin     []market.Candlestick
	BuySellTick    []BuySell
}

func (s *Stock) DrawVMin(count int) {

}

func (s *Stock) DrawVVVMin(count int) {

}

func (s *Stock) DrawDay(count int) {

}

func (s *Stock) DrawBuySellBalance() {

}
