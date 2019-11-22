package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type SuccessController struct {
	beego.Controller
}

func (this *SuccessController) Get() {
	this.Layout = "success.tpl"
	this.TplName = "success.tpl"
	fmt.Println("Success")

	this.Ctx.Redirect(302, "/download")
}
