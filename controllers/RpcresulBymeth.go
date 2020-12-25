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

func (r *ResultbymedController) Get() {
	var prejson models.Prejson
	err := r.ParseForm(&prejson)
	 fmt.Println(prejson)
	if err != nil {
		r.Ctx.WriteString("指令未接收")
		return
	}
	json, err := Bitcoin_conne.PrepareJSON(prejson.Method,prejson )
	if err != nil {
		r.Ctx.WriteString("指令处理出错")
		return
	}
	fmt.Println(json)

	data:=bitcoinService.Order(prejson.Method,json)

	fmt.Println("22",data)
	r.Data["A"]=data
	r.TplName = "result.html"
	/*
	for k, v =range map{
		if prejson.Args1 == k{
			return  v
	}


	}

	 */
}
