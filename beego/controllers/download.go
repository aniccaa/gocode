package controllers

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (this *DownloadController) Get() {
	this.Layout = "download.tpl"
	this.TplName = "download.tpl"

	fi, err := ioutil.ReadDir("static/download")
	if err != nil {
		fmt.Println("read dir error.")
		return
	}
	for _, f := range fi {
		if f.IsDir() {
			continue
		}
		fmt.Println(f.Name() + ": " + f.ModTime().String() + ": ")
	}

	// download the result file
	this.Ctx.Output.Download("static/download/result.xlsx", "result.xlsx")

}
