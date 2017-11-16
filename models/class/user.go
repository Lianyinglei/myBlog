package class

import
(
	"fmt"
	"github.com/astaxie/beego/orm"
	_"time"
)
type User struct {

	Id string `orm:"pk"`
	Nick string
	Email string
	Name string
	Phone string
	Password string
	Regtime string
	Private int



}

type Book struct {
	Id string `orm:"pk"`
	Content string
	Auther string

}

func TestORM()  {

	fmt.Println("hello")

	u := User{
		Id:       "lian",
		Nick:     "codeLian",
		Email:    "123@q.com",
		Name:"code",
		Phone:"18612341234",

	}

	o := orm.NewOrm()
	_, err := o.Insert(&u)

	u1 := User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	u2 := User{Nick: "jike"}
	err = o.Read(&u2, "nick")
	fmt.Println(u2)

	u2.Nick = "jike2"
	_, err = o.Update(&u2)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	//_, err = o.Delete(&u)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	if err != nil {
		fmt.Println(err)
	}

	u3 := User{Id:"lian"}

	err = o.Read(&u3)

	fmt.Println(u3)
}

func TestOrm()  {

	fmt.Println("hello testOrm")

}

func (u User) ReadDB() {

	o :=orm.NewOrm()
	err:= o.Read(&u)
	fmt.Println(err)

	return
}

func (u User) Create() {

	o:=orm.NewOrm()
	_,err:= o.Insert(&u)

	fmt.Println(err)
	return
}

func (u User)Updata()  {

	o:=orm.NewOrm()
	_,err:=o.Update(&u)

	fmt.Println(err)
	return
}

func (u User)Delete()  {

	o:=orm.NewOrm()
	_,err:= o.Delete(&u)
	fmt.Println(err)
	return
}