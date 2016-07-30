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
	c.Layout = "common/basic-layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	c.Data["contro_name"] = "login"
}

func (c *LoginController) Post() {

}
