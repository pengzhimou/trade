package utils

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
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
