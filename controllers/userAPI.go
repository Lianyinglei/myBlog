package controllers

import (
	"fmt"

	orm2 "github.com/astaxie/beego/orm"
	"myBlog/models/class"
	"github.com/astaxie/beego/validation"

	_"time"

)

func (c *UserController) Api_Profile() {

	type user struct {
		UserId string

		Hobby [] string

	}

	u:= user{
		"codeLian",
		[]string{"chess","code"},
	}

	c.Data["json"] = u

	c.ServeJSON()


}

func (c *UserController) Password(){

	fmt.Println("post 请求")
	type user struct {
		UserId string

		Hobby [] string
		Password string

	}

	u:= user{
		"codeLian--post",
		[]string{"postchess","codelian"},
		"password",
	}

	c.Data["json"] = u

	c.ServeJSON()


}


type RET struct {
	OK bool `json:"success"`
	Content interface{} `json:"content"`
}


func (c *UserController) Login()  {


	ret := RET{

		OK:true,
		Content:"profile.html",

	}

	defer func() {

		c.Data["json"] = ret
		c.ServeJSON()
	}()


}

func (c *UserController) ApiRegister(){

	o := orm2.NewOrm()
	u := class.User{Id:"lian"}
	err := o.Read(&u)
	fmt.Println(err)

	ret := RET{
		OK:true,
		Content:u,

	}

	c.Data["json"] =ret

	c.ServeJSON()

	//验证合法性

	id := c.GetString("userid")

	nick := c.GetString("nick")

	pwd1 := c.GetString("password")

	pwd2 := c.GetString("password2")

	email:=c.GetString("email")

	if len(nick)<1 {

		nick = id

	}


	valid := validation.Validation{}
	valid.Email(email,"Email")
	valid.Required(id, "Userid")
	valid.Required(pwd1, "pwd1")
	valid.Required(pwd2, "pwd2")

	valid.MaxSize(id, 20, "UserID")
	valid.MaxSize(nick, 30, "Nick")

	switch {

	case valid.HasErrors():

	case pwd1 != pwd2:
		valid.Error("两次密码不一致")

	default:
		//u:= class.User{
		//	Id:       id,
		//	Email:    email,
		//	Nick:     nick,
		//	Password: pwd1,
		//	Regtime:  time.Now(),
		//	Private:  10,
		//
		//}

		//u := &User{
		//	Id:       id,
		//	Email:    email,
		//	Nick:     nick,
		//	Password: PwGen(pwd1),
		//	Regtime:  time.Now(),
		//	Private:  DefaultPvt,
		//}

		//switch {
		//case u.
		//	valid.Error("用户名被占用")
		//case u.ExistEmail():
		//	valid.Error("邮箱被占用")
		//default:
		//	err := u.Create()
		//	if err == nil {
		//		return
		//	}
		//	valid.Error(Sprintf("%v", err))
		//}

	}

	//ret.Ok = false
	//ret.Content = valid.Errors[0].Key + valid.Errors[0].Message


	return


	//c.TplName = "user/register.html"

}

