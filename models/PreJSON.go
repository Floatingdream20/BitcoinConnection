package models

type Prejson struct {
	Method string      `form:"method"`
	Arg1    interface{} `form:"args1"`
	Arg2   interface{} `form:"args2"`
}
