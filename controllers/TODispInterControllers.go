package controllers

import (
	"Bitcoin_core_web/models"
	"github.com/astaxie/beego"
)

type Todisplnter struct {
	beego.Controller
}

func ( t *Todisplnter)Post(){
	var user models.User
     err:=t.ParseForm(&user)


	if err != nil {
		t.Ctx.WriteString("请输入账号密码")
		return
	}

      u,err:= user.Quereuser()

	if err != nil {
		t.TplName = "regist.html"
		return
	}

	t.Data["u"]=u.UserName
	t.TplName = "loginDis.html"

}
