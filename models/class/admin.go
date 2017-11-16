package class

import "fmt"

func TestAdmin()  {

	defer func () {

		fmt.Println("延迟调用的函数")
	}()

	fmt.Println("管理员权限")

	btnClick()
}

func btnClick()  {

	fmt.Println("调用内部函数")
}

