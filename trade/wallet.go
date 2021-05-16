package trade

import "github.com/huobirdcenter/huobi_golang/pkg/client"

//不同的交易中心
type StockWallet interface {
	Buy()
	Sell()
	ForceSell()
	ForceBuy()
}

type HuobiWallet struct {
	Client *client.WalletClient
}

func (hw *HuobiWallet) Buy(bStock string, bPrice, bValue float64) (dealstock string, dealPrice, dealValue float64, dealResult bool) {
	//挂单
	//检查成交
	return dealstock, dealPrice, dealValue, dealResult
}

func (hw *HuobiWallet) Sell() {

}

//止赢
func (hw *HuobiWallet) ForceBuy() {

}

//止损
func (hw *HuobiWallet) ForceCell() {

}
