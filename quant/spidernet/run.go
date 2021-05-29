package spidernet

import (
	"fmt"
	"trade/config"
	q "trade/quant"
	. "trade/utils"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

var (
	huobisystem = q.HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
	huobimarket = q.HuobiMarket{
		Client: new(client.MarketClient).Init(config.Host),
	}
	huobialgo = q.HuobiAlgo{
		Client: new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host),
	}
)

func Run() {
	ss := SpiderNet{}
	stocks := ss.GetAllAvaliableStock(&huobisystem, &huobimarket)
	for stock := range stocks {
		stk := q.Stock{
			Name: stock,
		}
		fmt.Println(ss.GetStockRiseFall(stk, &huobimarket, market.DAY1))
	}

}
