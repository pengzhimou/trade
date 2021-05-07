package rsi

type RSI struct {
	Code  string  `json:"code"`
	Date  string  `json:"date"` //2020-05-07
	Rsi6  float32 `json:"rsi6"`
	Rsi12 float32 `json:"rsi12"`
	Rsi24 float32 `json:"rsi24"`
}

type RSIFull struct {
	Code       string             `json:"code"`
	List       []RSI              `json:"list"`
	RealData   map[string]float32 `json:"realData"`
	Remark     string             `json:"remark"`
	ReturnCode int                `json:"ret_code"`
}

type RSIResp struct {
	ResBody  RSIFull `json:"showapi_res_body"`
	ResCode  int     `json:"showapi_res_code"`
	ResError string  `json:"showapi_res_error"`
	ResID    string  `json:"showapi_res_id"`
}
