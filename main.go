package main

import (
	"Bitcoin_core_web/dbMysqlSV"
	_ "Bitcoin_core_web/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	dbMysqlSV.ConDB()
	fmt.Println("hello world")
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/images", "./static/images")
	beego.Run()
}
