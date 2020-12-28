package models

/*
**JSON-RPC协议
**响应协议格式
 */
type Result struct {
	Id        int64         `json:"id"`
	Error     string         `json:"error"`
	Resultbit interface{} `json:"result"`//返回结果
}
