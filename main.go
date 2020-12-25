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
}
