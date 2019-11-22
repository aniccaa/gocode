package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type PostLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "login-page.tpl"
	this.TplName = "login-page.tpl"
}

func (this *MainController) Post() {
	this.Layout = "login-page.tpl"
	this.TplName = "login-page.tpl"

	login := PostLoginRequest{}
	if err := this.ParseForm(&login); err != nil {
		fmt.Println("handle form error.")
	}
	//fmt.Println(login.Username)
	//fmt.Println(login.Password)
	if login.Username == "admin" && login.Password == "admin123" {
		fmt.Println("login success!")
		this.Ctx.Redirect(302, "/excel")
	}
}
