package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/Go-SQL-Driver/MySQL"
	"myBlog/models/class"

	//"fmt"
	"fmt"
)

func init() {

	fmt.Println("链接数据库")
	//注册数据库驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//链接数据库
	orm.RegisterDataBase("default","mysql" ,"root:123456@tcp(localhost:3306)/jblog?charset=utf8")

//注册对象模型
	orm.RegisterModel(new(class.User))

	//开启同步
	orm.RunSyncdb("default",false,true)

	fmt.Println("链接成功")



}