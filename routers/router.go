package routers

import (
	"github.com/cho61183/babytime/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
