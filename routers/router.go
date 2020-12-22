package routers

import (
	"BitcoinConnection/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/toDisplayInterface",&controllers.Todisplnter{})
    beego.Router("/regist.html",&controllers.ToregController{})
    beego.Router("/givedbuserdata",&controllers.ToregController{})
    beego.Router("/RpcresultbyMethod",&controllers.ResultbymedController{})
}
