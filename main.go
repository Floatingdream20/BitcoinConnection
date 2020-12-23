package main

import (
	"Bitcoin_core_web/bitcoinService"
	"Bitcoin_core_web/dbMysqlSV"
	_ "Bitcoin_core_web/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	dbMysqlSV.ConDB()

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	//count,err:=bitcoinService.GetBlockCount()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(count)
	block,err:=bitcoinService.GetBlock("getblock","000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(block)
	return
	beego.Run()
}
