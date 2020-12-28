package models

/*
**JSON-RPC协议
**请求协议格式
 */
type RPCrequest struct {
	Id      int64       `json:"id"`//唯一标识
	Method  string      `json:"method"`//要调用的方法
	Jsonrpc string      `json:"jsonrpc"`//
	Params  interface{} `json:"params"`//调用时传的参数 数组【】格式
}
