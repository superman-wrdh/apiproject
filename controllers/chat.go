package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type ChatController struct {
	beego.Controller
}

func (c *ChatController) Finished() {
	defer func() {
		if err := recover(); err != nil {
			c.Abort("")
		}
	}()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 404 :uid is empty
// @router /:uid [get]
func (c *ChatController) Get() {
	a := 1
	x := 0
	a = a / x
	authString := c.Ctx.Input.Header("Authorization")
	uid := c.GetString(":uid")
	c.Data["json"] = map[string]interface{}{
		"uid":  uid,
		"time": time.Now(),
		"auth": authString,
	}
	c.ServeJSON()
}
