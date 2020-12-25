package controllers

import (
	"Bitcoin_core_web/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ToregController struct {
	beego.Controller
}
 func(t *ToregController)Get(){
	t.TplName = "regist.html"
}

func (t *ToregController)Post(){
	 var user models.User
	 err:=t.ParseForm(&user)
	 fmt.Println(user)
	if err != nil {
		t.Ctx.WriteString("你还有未填写的数据")
		return
	}
    _,err=user.Inseruser()
	if err != nil {
		t.Ctx.WriteString("用户数据未提交成功")
	}

  t.TplName = "index.html"

}