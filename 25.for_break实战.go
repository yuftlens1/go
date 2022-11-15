package main

import (
	"fmt"
)

//rand.Intn () 函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (), 并且传入一个变化的值作为参数，如 time.Now().UnixNano() , 就是可以生成时刻变化的值

// break 跳出整个循环体
func main() {
	/*
		for true {
			rand.Seed(time.Now().Unix())
			n := rand.Int() % 10 //生成一个10内的随机数!
			fmt.Println(n)
			if n == 9 || n == 2 || n == 3 || n == 4 {
				break // break 跳出整个for循环体
			}
		}

		   AK:
		   	for i := 0; i < 5; i++ {
		   		for j := 0; j < 9; j++ {
		   			if j > 5 {
		   				break AK //不加只跳出了内层j的for循环 //加了AK 就从前面AK i 外层停止循环
		   			}
		   			fmt.Printf("*")
		   		}
		   		fmt.Println("*")
		   	}
		   	fmt.Println()

	*/

	//项目登录实战
	var user, password string
	logintime := 3
	for true {
		for i := 0; i < logintime; i++ {
			fmt.Scanf("%s%s", &user, &password) //输入两份数据，输入的时候隔个空格就可以了
			if user == "yuft" && password == "test" {
				fmt.Println("登录成功")
				break
			} else {
				logintime-- //每次减一
				fmt.Printf("登录失败,你还有%d次机会\n", logintime)
			}
		}

	}

}
