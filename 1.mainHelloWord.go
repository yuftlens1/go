package main //包的申明   //go mod是go语言管理包的工具!

import "fmt" //引入包;导入输出库

func init() {
	// qaq := "测试"
	fmt.Println("init函数用于包(package)的初始化\n永远先执行init函数\n可以有多个init函数,main函数只能有一个\n")

}

func main() { //名字是main的函数 ()输入函数       {}输出函数
	fmt.Print("学习")
	fmt.Println("学习")
	fmt.Println("haha")
}

/*
Print() 函数不换行;
Println() 函数换行输出.Scanln应该同理
Printf() 格式化输出,Scanf同理
*/
