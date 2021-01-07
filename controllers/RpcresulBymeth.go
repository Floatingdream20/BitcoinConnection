package controllers

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"Bitcoin_core_web/bitcoinService"
	"Bitcoin_core_web/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ResultbymedController struct {
	beego.Controller
}
var Prejson models.Prejson
func (r *ResultbymedController) Get() {
	err := r.ParseForm(&Prejson)
	 fmt.Println(Prejson)
	if err != nil {
		r.Ctx.WriteString("指令未接收")
		return
	}
	jsons, err := Bitcoin_conne.PrepareJSON(Prejson.Method,Prejson )
	if err != nil {
		r.Ctx.WriteString("指令处理出错")
		return
	}
  fmt.Println(jsons)
	data:=bitcoinService.Order(Prejson.Method,jsons)


	fmt.Println(*data)
	r.Data["A"]=data
    r.TplName="loginDis.html"
	/*
	for k, v =range map{
		if prejson.Args1 == k{
			return  v
	}


	}

	 */
}
