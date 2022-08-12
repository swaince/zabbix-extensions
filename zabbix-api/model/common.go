package model

type ZabbixRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth"`
	Id      int         `json:"id"`
}

type ZabbixResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   Error       `json:"error,omitempty"`
	Id      int         `json:"id"`
}

type Error struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func NewRequest(method string, params interface{}) *ZabbixRequest {
	r := &ZabbixRequest{
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
	}

	return r
}

func NewResponse(result interface{}) *ZabbixResponse {
	r := &ZabbixResponse{
		Result: result,
	}
	return r
}
