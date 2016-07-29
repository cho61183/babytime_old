/*
	登录页面的控制

	@author liuyufeng
	@version 1.0
	@package controllers
*/
package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
