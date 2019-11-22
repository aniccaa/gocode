package router

import (
	"awesomeProject/beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/excel", &controllers.ExcelController{})
	beego.Router("/download", &controllers.DownloadController{})
	beego.Router("/success", &controllers.SuccessController{})
}
