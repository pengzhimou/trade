package main

import (
	"trade/config"
	"trade/quant/spidernet"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
)

func placeOrder() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "adausdt",
		Price:     "1.5",
		Amount:    "5", //币的数量
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Place order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func main() {
	// example.W()
	// example.X()
	// example.Y()
	// example.Z()

	// placeOrder()
	spidernet.Run()

}
