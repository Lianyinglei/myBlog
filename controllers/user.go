package controllers

import (

	_"github.com/astaxie/beego"
	"github.com/astaxie/beego"
)


type UserController struct {

	beego.Controller
	//BaseController
}

func (c *UserController) Profile(){

	c.TplName = "user/profile.html"
}


func (c *UserController) Put()  {

	c.Ctx.WriteString("hello lian")
}

func (c *UserController) Join() {

	c.TplName = "user/join.html"
}

func (c *UserController) Register(){

	c.TplName = "user/register.html"

}


func (c *UserController)  Setting(){

	c.TplName = "user/setting.html"

}




