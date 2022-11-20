package main

import (
	"fmt"
	"time"
)

func main() {
	//项目登录实战
	var user, password string
	var logintime = 6
AK:
	for {
		for ; logintime > 0; logintime = logintime - 1 { //for i := 0; i < logintime; i++ {
			fmt.Println("请输入用户名密码")
			fmt.Scanf("%s%s\n", &user, &password) //输入两份数据，输入的时候隔个空格就可以了  !!!%s%s\n  win上必须加\n!!!天坑记录下
			/*
				Windows默认换行符是  \r\n    !!!
				我在windows上测试过。猜测你可能是输入1,2,3加上了回车，1,2,3被第一次循环的fmt.Scanf读取了，第二次循环，fmt.Scanf读到你的enter就认为输入结束了，所以输出了两次1 2 3。有两种解决方案：
				输入 1,2,3 4,5,5 再回车，这样会输出两行，一行1 2 3，一行4 5 6
				2.将fmt.Scanf格式符改为fmt.Scanf("%d,%d,%d\n", &a,&b,&c)即将格式符后面加上\n，强制要求每次输入带上回车。
			*/
			if user == "yuft" && password == "test" {
				fmt.Println("登录成功")
				break AK //这里的break跳出了里面这个for循环
			} else {
				fmt.Printf("登录失败,你还有%d次机会\n", logintime-1)
			}
		}
		fmt.Print("机会用完，请30分钟后再试\n")
		time.Sleep(time.Second * 30)
		logintime = 6
	}
}

/*
// rand.Intn () 函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (), 并且传入一个变化的值作为参数，如 time.Now().UnixNano() , 就是可以生成时刻变化的值
// break 跳出整个循环体
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
