package main

import (
	"Bitcoin_core_web/bitcoinService"
	_ "Bitcoin_core_web/routers"
)

func main() {
	//dbMysqlSV.ConDB()
	//
	//beego.SetStaticPath("/js", "./static/js")
	//beego.SetStaticPath("/css", "./static/css")
	//beego.SetStaticPath("/img", "./static/img")
	//
	//beego.Run()
	bitcoinService.GetMemoryInfo()
}

//整理运行思路
//尝试
//RPCutilt strivice RpcresulBymeth
