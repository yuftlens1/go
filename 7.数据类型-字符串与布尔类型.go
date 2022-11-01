package main

import (
	"fmt"
)

// 布尔值
func mainn() {
	var isok bool = false
	if isok {
		fmt.Println("执行a计划。")
	} else {
		fmt.Println("a计划失败,启动b计划。")
	}
}

// 字符串;字符串变量的只要用双引号引起来
func main() {
	var mingling string = "calc"
	fmt.Println(mingling)
	//exec.Command(mingling).Run() //执行系统命令
	mingling = "notepad" //字符串为可修改数据类型。上面已经定义mingling这个变量了，所以不需要写var string
	fmt.Println("ming\nling_hahahahah")
	//exec.Command(mingling).Run() //执行系统命令
}
