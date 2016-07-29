/*
注册页面的控制层

@author liuyufeng
@version 1.0
@package controllers
*/
package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.tpl"
}
