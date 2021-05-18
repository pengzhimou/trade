package quant

import (
	"strconv"
	"trade/config"
	"trade/utils"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/algoorder"
)

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
func (ha *HuobiAlgo) OrderCreateOld(stock, limitmarket, buysell, amount, startprice, stopprice string) (algoorder.PlaceOrderRequest, *algoorder.PlaceOrderResponse, error) {

	accountId, _ := strconv.Atoi(config.AccountId)
	request := algoorder.PlaceOrderRequest{
		AccountId:     accountId,
		Symbol:        stock,       // adausdt
		OrderType:     limitmarket, // limit/market
		OrderSize:     amount,      // "6"
		OrderPrice:    startprice,  // "2.0"
		OrderSide:     buysell,     // buy/sell
		TimeInForce:   "gtc",
		ClientOrderId: utils.GenUUID(),
		StopPrice:     stopprice, // stopprice
		// TrailingRate:  "",
	}
	resp, err := ha.Client.PlaceOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			applogger.Info("Place algo order successfully, client order id: %s", resp.Data.ClientOrderId)
		} else {
			applogger.Error("Place algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
	return request, resp, err
}

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
