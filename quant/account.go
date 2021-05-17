package quant

import (
	"strconv"
	"trade/config"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
)

func GetAccountInfo() ([]account.AccountInfo, error) {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
	}
	return resp, err
}

func GetSpotAccountID() string {
	resp, err := GetAccountInfo()
	if err != nil {
		return config.AccountId
	}
	for _, act := range resp {
		if act.Type == "spot" {
			return strconv.FormatInt(act.Id, 10)
		}
	}
	return config.AccountId
}
