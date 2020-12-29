package models

type Prejson struct {
	Method string      `form:"method"`
	Args    interface{} `form:"args"`
}
