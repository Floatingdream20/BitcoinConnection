package controllers

import (
	"BitcoinConnection/Bitcoin_conne"
	"BitcoinConnection/bitcoin_server"
	"BitcoinConnection/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"strings"
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

	GetBlockCount:=bitcoin_server.GetBlockCount()
    fmt.Println(GetBlockCount)

	body, err := Bitcoin_conne.PrepareJSON(prejson.Method, prejson.Args)
	if err != nil {
		r.Ctx.WriteString("初始结构json化出错")
	}
	Rpcresult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(body))
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

}
