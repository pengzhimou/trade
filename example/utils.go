package example

import (
	"encoding/json"
	"fmt"
)

func P(s interface{}) string {
	aaa, _ := json.Marshal(s)
	rst := string(aaa)
	fmt.Println(rst)
	return (rst)
}
