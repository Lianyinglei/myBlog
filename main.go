package main

import (
	_ "myBlog/routers"
	"github.com/astaxie/beego"
	_ "myBlog/models"
	"myBlog/models/class"


)

func main() {

	class.TestORM()
	//class.TestOrm()
   class.TestAdmin()

//u:= class.User{
//	Name:"lian",
//	Phone:"186109403380",
//	Id:"codeocode",
//	Nick:"haha",
//	Regtime:"2017-10-11",
//
//}
//
//   u.Create()


	beego.Run()

}

