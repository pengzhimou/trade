package example

import (
	"fmt"
	"trade/config"
	"trade/quant"
	"trade/utils"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/algoorder"
)

var (
	huobialgo = quant.HuobiAlgo{
		Client: new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host),
	}
)

func Z() {
	// Trade_Mix()
	Trade_OrderCreate()
	// Trade_OrdersGetOpen()
	// Trade_OrdersHistory()
	// Trade_OrdersGet()
	// Trade_OderCancle()
}

func Trade_OrderCreate() {
	fmt.Println("@@@@@@@@@@Trade_OrderCreate")
	orderRequest := algoorder.PlaceOrderRequest{
		Symbol:        "adausdt", // adausdt
		OrderType:     "limit",   // limit/market
		OrderSize:     "5",       // "6"
		OrderPrice:    "1.5",     // "2.0"
		OrderSide:     "buy",     // buy/sell
		TimeInForce:   "gtc",
		ClientOrderId: utils.GenUUID(),
		StopPrice:     "1.4", // stopprice
		// TrailingRate:  "",
	}

	aaa, bbb, _ := huobialgo.OrderCreate(orderRequest)
	utils.P(aaa)
	utils.P(bbb)
}

func Trade_OrdersGetOpen() {
	fmt.Println("@@@@@@@@@@Trade_OrdersGetOpen")
	aaa, _ := huobialgo.OrdersGetOpen()
	utils.P(aaa)
}

func Trade_OrdersHistory() {
	fmt.Println("@@@@@@@@@@Trade_OrdersHistory")
	aaa, _ := huobialgo.OrdersHistory("adausdt")
	utils.P(aaa)
}

func Trade_OrdersGet() {
	fmt.Println("@@@@@@@@@@Trade_OrdersGet")
	x, _ := huobialgo.OrdersGetOpen()

	aaa, _ := huobialgo.OrderGet(x.Data[0].ClientOrderId)
	utils.P(aaa)
}

func Trade_OderCancle() {
	fmt.Println("@@@@@@@@@@Trade_OderCancle")
	x, _ := huobialgo.OrdersGetOpen()

	for _, od := range x.Data {
		fmt.Println("========")
		aaa, _ := huobialgo.OderCancle(od.ClientOrderId)
		utils.P(aaa)
	}
}

func Trade_Mix() {

	// stock := "adausdt"
	// buyprice := "1.96"
	// fmt.Println("@@@@@@@@@@Trade_buy")
	// aaa, bbb, _ := huobialgo.OrderCreate(stock, "limit", "buy", "10", buyprice, buyprice)
	// utils.P(aaa)
	// utils.P(bbb)

	// stock := "xmrusdt"
	// sellprice := "337"
	// fmt.Println("@@@@@@@@@@Trade_cel")
	// aaa, bbb, _ := huobialgo.OrderCreate(stock, "limit", "sell", "5", sellprice, sellprice)
	// utils.P(aaa)
	// utils.P(bbb)

	xxx, _ := huobialgo.OderCancle("0bd336d5-2d0e-4283-938a-7e1b9a3e48c7")
	utils.P(xxx)

}
