package controllers
import (
	"github.com/astaxie/beego"
	"myProfile/models/class"

)


type BaseController struct {

	beego.Controller
}

func (c *BaseController) DoLogin(u class.User)  {

	c.SetSession("user",u)

}

func (c *BaseController) DoLogout()  {

	c.DestroySession()
	c.Abort("302")
}

//检查登录
func (c *BaseController) IsLogin() bool  {

	return c.GetSession("user") != nil
}

//检查用户是否登录

func (c *BaseController) CheckLoign()  {

	if !c.IsLogin() {
		c.Redirect("json",302)

		c.Abort("302")
	}
}


func (c *BaseController) Prepare()  {
	if c.IsLogin() {

		c.Data["user"] = c.GetSession("user").(class.User)
	}
}

