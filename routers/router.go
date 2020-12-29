package routers

import (
	"Bitcoin_core_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index.html",&controllers.MainController{})
    beego.Router("/toDisplayInterface",&controllers.Todisplnter{})
    beego.Router("/register.html",&controllers.ToregController{})
    beego.Router("/dbUserData",&controllers.ToregController{})
    beego.Router("/rpcResultByMethod",&controllers.ResultbymedController{})
}
