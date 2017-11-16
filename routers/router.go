package routers

import (
	"myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})
    beego.Router("user/profile",&controllers.UserController{},`get:Profile`)
	beego.Router("/api/profile",&controllers.UserController{},`get:Api_Profile`)
	beego.Router("user/password",&controllers.UserController{},`post:Password`)
	beego.Router("/user/login",&controllers.UserController{},`get:Join`)

    beego.Router("/user/put",&controllers.UserController{},`get:Put`)
	beego.Router("/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/register",&controllers.UserController{},`get:Register`)
	beego.Router("/user/register",&controllers.UserController{},`post:ApiRegister`)
	beego.Router("user/setting",&controllers.UserController{},`get:Setting`)



}
