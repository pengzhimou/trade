package utils

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func P(s interface{}) string {
	aaa, _ := json.Marshal(s)
	rst := string(aaa)
	fmt.Println(rst)
	return (rst)
}

func GenUUID() string {
	u2 := uuid.NewV4()
	return u2.String()
}

func Cal(a, b decimal.Decimal, w string) decimal.Decimal {
	switch w {
	case "+":
		return a.Add(b)
	case "-":
		return a.Sub(b)
	case "*":
		return a.Mul(b)
	case "/":
		return a.Div(b)
	default:
		panic("+-*/")
	}
}

func Blacklist() []string {
	return []string{
		// "ftiusdt",
		// "swftcusdt",
	}
}
