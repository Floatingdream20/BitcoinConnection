package controllers

import (
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
	if err != nil {
		r.Ctx.WriteString("指令未接收")
		return
	}
	data,err:=bitcoinService.Order(prejson.Method,prejson.Args1)
	fmt.Println("22",data)
	r.Data["A"]=data
	r.TplName = "result.html"
	/*
	count,err:=bitcoinService.GetBlockCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
	r.TplName = "result.html"

	//,_ := Bitcoin_conne.PrepareJSON(prejson.Method, prejson.Args)
	if err != nil {
		r.Ctx.WriteString("初始结构json化出错")
	}
	Rpcresult, err := Bitcoin_conne.Dopost("", nil, "")
	fmt.Println(Rpcresult)
	if err != nil {
		r.Ctx.WriteString("数据处理出错")
	}
	if Rpcresult == nil {
		r.Ctx.WriteString("结果为空")
		return
	}
	A := &(Rpcresult.Data).Resultbit
	fmt.Println(*A)
	if Rpcresult.Code == http.StatusOK {
		r.Data["A"] = A
		r.TplName = "result.html"

	} else {
		fmt.Println("失败")

		r.Ctx.WriteString("获取失败")
		return
	}
*/
}
