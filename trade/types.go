package trade

// type GetCandlestickResponse struct {
// 	Status string        `json:"status"`
// 	Ch     string        `json:"ch"`
// 	Ts     int64         `json:"ts"`
// 	Data   []Candlestick `json:"data"`
// }

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

// type GetLast24hCandlestickAskBidResponse struct {
// 	Status string             `json:"status"`
// 	Ch     string             `json:"ch"`
// 	Ts     int64              `json:"ts"`
// 	Tick   *CandlestickAskBid `json:"tick"`
// }

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

// type GetAllSymbolsLast24hCandlesticksAskBidResponse struct {
// 	Status string              `json:"status"`
// 	Ts     int64               `json:"ts"`
// 	Data   []SymbolCandlestick `json:"data"`
// }

// type SymbolCandlestick struct {
// 	Amount  decimal.Decimal `json:"amount"`
// 	Open    decimal.Decimal `json:"open"`
// 	Close   decimal.Decimal `json:"close"`
// 	High    decimal.Decimal `json:"high"`
// 	Symbol  string          `json:"symbol"`
// 	Count   int64           `json:"count"`
// 	Low     decimal.Decimal `json:"low"`
// 	Vol     decimal.Decimal `json:"vol"`
// 	Bid     decimal.Decimal `json:"bid"`
// 	BidSize decimal.Decimal `json:"bidSize"`
// 	Ask     decimal.Decimal `json:"ask"`
// 	AskSize decimal.Decimal `json:"askSize"`
// }

// type GetDepthResponse struct {
// 	Status string `json:"status"`
// 	Ch     string `json:"ch"`
// 	Ts     int64  `json:"ts"`
// 	Tick   *Depth `json:"tick"`
// }

// type Depth struct {
// 	Timestamp int64               `json:"ts"`
// 	Version   int64               `json:"version"`
// 	Bids      [][]decimal.Decimal `json:"bids"`
// 	Asks      [][]decimal.Decimal `json:"asks"`
// }

// type GetLatestTradeResponse struct {
// 	Status string     `json:"status"`
// 	Ch     string     `json:"ch"`
// 	Ts     int64      `json:"ts"`
// 	Tick   *TradeTick `json:"tick"`
// }

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
