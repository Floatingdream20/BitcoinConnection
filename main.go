package main

import (
	"Bitcoin_core_web/dbMysqlSV"
	_ "Bitcoin_core_web/routers"
	"github.com/astaxie/beego"
)

func main() {
	dbMysqlSV.ConDB()

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()

	//json, err := Bitcoin_conne.PrepareJSON("getmemoryinfo",controllers.Prejson)
	//if err != nil {
	//	fmt.Println("有问题")
	//	return
	//}
	//fmt.Println(json)
	//bitcoinService.GetMemoryInfo(json)
	//bitcoinService.Txoutset(json)
	//bitcoinService.MemPoolInfo(json)
	//bitcoinService.ChainTxStats(json)
	//bitcoinService.BlockHeader(json)
}

//整理运行思路
//尝试
//RPCutilt strivice RpcresulBymeth
