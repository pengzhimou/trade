package main

import (
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

// "github.com/huobirdcenter/huobi_golang/pkg/client"
// "github.com/huobirdcenter/huobi_golang/logging/applogger"
// "github.com/huobirdcenter/huobi_golang/pkg/client/orderwebsocketclient"
// "github.com/huobirdcenter/huobi_golang/pkg/model/auth"
// "github.com/huobirdcenter/huobi_golang/pkg/model/order"

func main() {

	// Get the list of accounts owned by this API user and print the detail on console
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
	} else {
		applogger.Info("Get account, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("account: %+v", result)
		}
	}

	// // client := new(orderwebsocketclient.RequestOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	// // Initialize a new instance
	// client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// // Set the callback handlers
	// client.SetHandler(
	// 	// Connected handler
	// 	func(resp *auth.WebSocketV2AuthenticationResponse) {
	// 		if resp.IsSuccess() {
	// 			// Subscribe if authentication passed
	// 			client.Subscribe("btcusdt", "1149")
	// 		} else {
	// 			applogger.Info("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
	// 		}
	// 	},
	// 	// Response handler
	// 	func(resp interface{}) {
	// 		subResponse, ok := resp.(order.SubscribeOrderV2Response)
	// 		if ok {
	// 			if subResponse.Action == "sub" {
	// 				if subResponse.IsSuccess() {
	// 					applogger.Info("Subscription topic %s successfully", subResponse.Ch)
	// 				} else {
	// 					applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
	// 				}
	// 			} else if subResponse.Action == "push" {
	// 				if subResponse.Data != nil {
	// 					o := subResponse.Data
	// 					applogger.Info("Order update, event: %s, symbol: %s, type: %s, status: %s",
	// 						o.EventType, o.Symbol, o.Type, o.OrderStatus)
	// 				}
	// 			}
	// 		} else {
	// 			applogger.Warn("Received unknown response: %v", resp)
	// 		}
	// 	})

	// // Connect to the server and wait for the handler to handle the response
	// client.Connect(true)

}
