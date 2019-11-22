package main

import (
	_ "awesomeProject/beego/router"

	"github.com/astaxie/beego"
)

func main() {
	// beego.BConfig.Listen.ListenTCP4 = true
	if err := beego.LoadAppConfig("ini", "beego/conf/app.conf"); err != nil {
		// dir, _ := os.Getwd()
		return
	}

	beego.Run(":8080")
}
