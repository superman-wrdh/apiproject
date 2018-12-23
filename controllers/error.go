package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = map[string]string{
		"code": "404",
	}
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = map[string]string{
		"code": "500",
	}
	c.ServeJSON()
}

func (c *ErrorController) Error600() {
	c.Data["json"] = map[string]string{
		"code": "600",
	}
	c.ServeJSON()
}

func (c *ErrorController) Error() {
	c.Data["json"] = map[string]string{
		"code": "服务器内部错误，请联系管理员",
	}
	c.ServeJSON()
}
