package trade

import (
	"fmt"
	"strconv"
	"trade/config"
	"trade/example"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/algoorder"
)

// {
// 	"account-id": "16936884",
// 	"symbol": "adausdt",
// 	"type": "buy-limit",
// 	"amount": "5",
// 	"price": "1.9",
// 	"source": "spot-api",
// 	"client-order-id": "bbe84447-7ef6-4ffa-9cfb-1b2ad0452175",
// 	"stop-price": "2.0",
// 	"operator": "gte"
//   }
func OrderCreate() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host)

	accountId, _ := strconv.Atoi(config.AccountId)
	fmt.Println(accountId)
	request := algoorder.PlaceOrderRequest{
		AccountId:     accountId,
		Symbol:        "adausdt",
		OrderPrice:    "2.0",
		OrderSide:     "buy",
		OrderSize:     "5",
		TimeInForce:   "gtc",
		OrderType:     "limit",
		ClientOrderId: "huobi1901",
		StopPrice:     "2.1",
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			applogger.Info("Place algo order successfully, client order id: %s", resp.Data.ClientOrderId)
		} else {
			applogger.Error("Place algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func OderCancle() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := algoorder.CancelOrdersRequest{
		ClientOrderIds: []string{"huobi1901"},
	}
	resp, err := client.CancelOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data.Accepted != nil {
				for _, id := range resp.Data.Accepted {
					applogger.Info("Cancelled client order id success: %s", id)
				}
			}
			if resp.Data.Rejected != nil {
				for _, id := range resp.Data.Rejected {
					applogger.Error("Cancelled client order id error: %s", id)
				}
			}
		} else {
			applogger.Error("Cancel algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func OrdersGetOpen() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", config.AccountId)
	fmt.Println(config.AccountId)

	resp, err := client.GetOpenOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				example.P(resp)
				applogger.Info("There are total %d open orders", len(resp.Data))
				for _, o := range resp.Data {
					applogger.Info("Open orders, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
				}
			}
		} else {
			applogger.Error("Get open order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func OrdersHistory() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "htusdt")
	request.AddParam("orderStatus", "canceled")

	resp, err := client.GetHistoryOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				applogger.Info("There are total %d history orders", len(resp.Data))
				for _, o := range resp.Data {
					applogger.Info("history orders, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
				}
			}
		} else {
			applogger.Error("Get history order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func OrderGet() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", "huobi1901")

	resp, err := client.GetSpecificOrder(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				o := resp.Data
				applogger.Info("Get order, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
			}
		} else {
			applogger.Error("Get order error, code: %s, message: %s", resp.Code, resp.Message)
		}
	}
}
