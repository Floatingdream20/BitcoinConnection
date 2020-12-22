package main

import (
	"Bitcoin_core_web/bitcoinService"
	_ "Bitcoin_core_web/routers"
	"fmt"
)

func main() {
//	dbMysqlSV.ConDB()
//
//beego.SetStaticPath("/js","./static/js")
//beego.SetStaticPath("/css","./static/css")
//beego.SetStaticPath("/img","./static/img")
	count,err:=bitcoinService.GetBlockCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
	//beego.Run()
}

