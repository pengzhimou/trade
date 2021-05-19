package quant

import (
	"strconv"
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/algoorder"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
)

type HuobiTrade struct {
	Client *client.OrderClient
}

func (ht *HuobiTrade) OrderCreate(orderRequest order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	resp, err := ht.Client.PlaceOrder(&orderRequest)
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
	return resp, err
}

func (ht *HuobiTrade) OrdersCreate(ordersRequest []order.PlaceOrderRequest) (*order.PlaceOrdersResponse, error) {
	resp, err := ht.Client.PlaceOrders(ordersRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					if r.OrderId != 0 {
						applogger.Info("Place order successfully: order id %d", r.OrderId)
					} else {
						applogger.Info("Place order error: %s", r.ErrorMessage)
					}
				}
			}
		case "error":
			applogger.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
	return resp, err
}

func (ht *HuobiTrade) OrderGetOpen(stock string) (*order.GetOpenOrdersResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("account-id", GetSpotAccountID())
	request.AddParam("symbol", stock)
	resp, err := ht.Client.GetOpenOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					applogger.Info("Open orders, symbol: %s, price: %s, amount: %s", o.Symbol, o.Price, o.Amount)
				}
				applogger.Info("There are total %d open orders", len(resp.Data))
			}
		case "error":
			applogger.Error("Get open order error: %s", resp.ErrorMessage)
		}
	}
	return resp, err
}

func (ht *HuobiTrade) OrderGetHistory(stock, state string) (*order.GetHistoryOrdersResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", stock)
	request.AddParam("states", state)
	resp, err := ht.Client.GetHistoryOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					applogger.Info("Order history, symbol: %s, price: %s, amount: %s, state: %s", o.Symbol, o.Price, o.Amount, o.State)
				}
				applogger.Info("There are total %d orders", len(resp.Data))
			}
		case "error":
			applogger.Error("Get history order error: %s", resp.ErrorMessage)
		}
	}
	return resp, err
}

func (ht *HuobiTrade) OrderGet(orderId string) (*order.GetOrderResponse, error) {
	resp, err := ht.Client.GetOrderById(orderId)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				applogger.Info("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			applogger.Error("Get order by id error: %s", resp.ErrorMessage)
		}
	}
	return resp, err
}

func (ht *HuobiTrade) OrderCancle(orderId string) (*order.CancelOrderByIdResponse, error) {
	resp, err := ht.Client.CancelOrderById(orderId)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Cancel order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
	return resp, err
}

// //??????????
// func (ht *HuobiTrade) OrderCancleByClient(orderId string) (*order.CancelOrderByClientResponse, error) {
// 	resp, err := ht.Client.CancelOrderByClientOrderId(orderId)
// 	if err != nil {
// 		applogger.Error(err.Error())
// 	} else {
// 		switch resp.Status {
// 		case "ok":
// 			applogger.Info("Cancel order successfully, order id: %d", resp.Data)
// 		case "error":
// 			applogger.Info("Cancel order error: %s", resp.ErrorMessage)
// 		}
// 	}
// 	return resp, err
// }

// // cancelRequest := order.CancelOrdersByCriteriaRequest{
// // 	AccountId: config.AccountId,
// // 	Symbol:    "btcusdt",
// // }
// func (ht *HuobiTrade) OrderCancelByCriteria(cancelRequest order.CancelOrdersByCriteriaRequest) (*order.CancelOrdersByCriteriaResponse, error) {
// 	resp, err := ht.Client.CancelOrdersByCriteria(&cancelRequest)
// 	if err != nil {
// 		applogger.Error(err.Error())
// 	} else {
// 		switch resp.Status {
// 		case "ok":
// 			if resp.Data != nil {
// 				d := resp.Data
// 				applogger.Info("Cancel orders successfully, success count: %d, failed count: %d, next id: %d", d.SuccessCount, d.FailedCount, d.NextId)
// 			}
// 		case "error":
// 			applogger.Error("Cancel orders error: %s", resp.ErrorMessage)
// 		}
// 	}
// 	return resp, err
// }

// @@@@@@@@@@Trade_OrderCreate
// 25997946
// 2021-05-18T00:53:03.057+0800	INFO	Place algo order successfully, client order id: e6fabadc-d3e1-483c-9de8-fd6df59df043
// {"accountId":25997946,"symbol":"adausdt","orderPrice":"1.9","orderSide":"buy","orderSize":"5","orderValue":"","timeInForce":"gtc","orderType":"limit","clientOrderId":"e6fabadc-d3e1-483c-9de8-fd6df59df043","stopPrice":"2.0","trailingRate":""}
// {"code":200,"message":"","Data":{"clientOrderId":"e6fabadc-d3e1-483c-9de8-fd6df59df043"}}
// @@@@@@@@@@Trade_OrdersGetOpen
// 2021-05-18T00:53:03.085+0800	INFO	There are total 1 open orders
// 2021-05-18T00:53:03.086+0800	INFO	Open orders, cid: e6fabadc-d3e1-483c-9de8-fd6df59df043, symbol: adausdt, status: created
// {"code":200,"message":"","Data":[{"accountId":25997946,"source":"api","clientOrderId":"e6fabadc-d3e1-483c-9de8-fd6df59df043","symbol":"adausdt","orderPrice":"1.9","orderSize":"5","orderValue":"","orderSide":"buy","timeInForce":"gtc","orderType":"limit","stopPrice":"2","trailingRate":"","orderOrigTime":1621270383050,"lastActTime":1621270383055,"orderStatus":"created"}],"nextId":0}
// @@@@@@@@@@Trade_OrdersHistory
// 2021-05-18T00:53:03.112+0800	INFO	There are total 0 history orders
// {"code":200,"message":"","Data":[],"nextId":0}
// @@@@@@@@@@Trade_OrdersGet
// 2021-05-18T00:53:03.132+0800	INFO	There are total 1 open orders
// 2021-05-18T00:53:03.132+0800	INFO	Open orders, cid: e6fabadc-d3e1-483c-9de8-fd6df59df043, symbol: adausdt, status: created
// 2021-05-18T00:53:03.159+0800	INFO	Get order, cid: e6fabadc-d3e1-483c-9de8-fd6df59df043, symbol: adausdt, status: created
// {"code":200,"message":"","Data":{"accountId":25997946,"source":"api","clientOrderId":"e6fabadc-d3e1-483c-9de8-fd6df59df043","orderId":"","symbol":"adausdt","orderPrice":"1.9","orderSize":"5","orderValue":"","orderSide":"buy","timeInForce":"gtc","orderType":"limit","stopPrice":"2","trailingRate":"","orderOrigTime":1621270383050,"lastActTime":1621270383091,"orderCreateTime":0,"orderStatus":"created","errCode":0,"errMessage":""}}
// @@@@@@@@@@Trade_OderCancle
// 2021-05-18T00:53:03.181+0800	INFO	There are total 1 open orders
// 2021-05-18T00:53:03.181+0800	INFO	Open orders, cid: e6fabadc-d3e1-483c-9de8-fd6df59df043, symbol: adausdt, status: created
// ========
// 2021-05-18T00:53:03.221+0800	INFO	Canselled client order id success: e6fabadc-d3e1-483c-9de8-fd6df59df043
// {"code":200,"message":"","Data":{"accepted":["e6fabadc-d3e1-483c-9de8-fd6df59df043"],"rejected":[]}}

type HuobiAlgo struct {
	Client *client.AlgoOrderClient
}

// type PlaceOrderRequest struct {
// 	AccountId     int    `json:"accountId"`
// 	Symbol        string `json:"symbol"`
// 	OrderPrice    string `json:"orderPrice"`
// 	OrderSide     string `json:"orderSide"`
// 	OrderSize     string `json:"orderSize"`
// 	OrderValue    string `json:"orderValue"`
// 	TimeInForce   string `json:"timeInForce"`
// 	OrderType     string `json:"orderType"`
// 	ClientOrderId string `json:"clientOrderId"`
// 	StopPrice     string `json:"stopPrice"`
// 	TrailingRate  string `json:"trailingRate"`
// }
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
func (ha *HuobiAlgo) OrderCreate(orderRequest algoorder.PlaceOrderRequest) (algoorder.PlaceOrderRequest, *algoorder.PlaceOrderResponse, error) {

	accountId, _ := strconv.Atoi(config.AccountId)
	orderRequest.AccountId = accountId
	resp, err := ha.Client.PlaceOrder(&orderRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			applogger.Info("Place algo order successfully, client order id: %s", resp.Data.ClientOrderId)
		} else {
			applogger.Error("Place algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
	return orderRequest, resp, err
}

// func placeOrder() {
// 	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
// 	request := order.PlaceOrderRequest{
// 		AccountId: config.AccountId,
// 		Type:      "buy-limit",
// 		Source:    "spot-api",
// 		Symbol:    "btcusdt",
// 		Price:     "1.1",
// 		Amount:    "1",
// 	}
// 	resp, err := client.PlaceOrder(&request)
// 	if err != nil {
// 		applogger.Error(err.Error())
// 	} else {
// 		switch resp.Status {
// 		case "ok":
// 			applogger.Info("Place order successfully, order id: %s", resp.Data)
// 		case "error":
// 			applogger.Error("Place order error: %s", resp.ErrorMessage)
// 		}
// 	}
// }

func (ha *HuobiAlgo) OrdersGetOpen() (*algoorder.GetOpenOrdersResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", config.AccountId)

	resp, err := ha.Client.GetOpenOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				// example.P(resp)
				applogger.Info("There are total %d open orders", len(resp.Data))
				for _, o := range resp.Data {
					applogger.Info("Open orders, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
				}
			}
		} else {
			applogger.Error("Get open order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
	return resp, err
}

func (ha *HuobiAlgo) OrdersHistory(stock string) (*algoorder.GetHistoryOrdersResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", stock)
	request.AddParam("orderStatus", "canceled")

	resp, err := ha.Client.GetHistoryOrders(request)
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
	return resp, err
}

// {
// 	"Data": {
// 	  "accountId": 25997946,
// 	  "clientOrderId": "a4d57783-4919-43ce-b4c4-f064e2fd2e30",
// 	  "errCode": 0,
// 	  "errMessage": "",
// 	  "lastActTime": "1621352033882",
// 	  "orderCreateTime": 0,
// 	  "orderId": "",
// 	  "orderOrigTime": "1621352033688",
// 	  "orderPrice": "1.5",
// 	  "orderSide": "buy",
// 	  "orderSize": "5",
// 	  "orderStatus": "created",
// 	  "orderType": "limit",
// 	  "orderValue": "",
// 	  "source": "api",
// 	  "stopPrice": "1.5",
// 	  "symbol": "adausdt",
// 	  "timeInForce": "gtc",
// 	  "trailingRate": ""
// 	},
// 	"code": 200,
// 	"message": ""
//   }
func (ha *HuobiAlgo) OrderGet(orderId string) (*algoorder.GetSpecificOrderResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", orderId)

	resp, err := ha.Client.GetSpecificOrder(request)
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
	return resp, err
}

func (ha *HuobiAlgo) OderCancle(orderId string) (*algoorder.CancelOrdersResponse, error) {
	request := algoorder.CancelOrdersRequest{
		ClientOrderIds: []string{orderId},
	}

	resp, err := ha.Client.CancelOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data.Accepted != nil {
				for _, id := range resp.Data.Accepted {
					applogger.Info("Canselled client order id success: %s", id)
				}
			}
			if resp.Data.Rejected != nil {
				for _, id := range resp.Data.Rejected {
					applogger.Error("Canselled client order id error: %s", id)
				}
			}
		} else {
			applogger.Error("Cancel algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
	return resp, err
}
