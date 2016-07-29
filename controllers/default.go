package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	nums := []int{1, 3, 4, 5, 6, 7, 8, 9, 0}
	c.Data["Nums"] = nums
	c.Data["Html"] = "<div> hello babytime </div>"
}
