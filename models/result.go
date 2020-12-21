package models

type Result struct {
	Id int64 `json:"id"`
	Error string `json:"error"`
	Resultbit interface{} `json:"result"`
}
