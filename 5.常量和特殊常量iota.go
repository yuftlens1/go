package main

import (
	"fmt"
	"os/exec"
)

func main() {
	const name string = "yft" //常量不管有没有被使用都不影响代码运行
	//var age int = 12   //变量被使用了程序才能跑，有未被使用的变量--程序不能跑！
	const 计算器, 记事本 = "calc", "notepad"
	fmt.Println(计算器)
	exec.Command(计算器).Run() //这里的计算器不能加引号，加了就是成字符串了

	const (
		练气 = iota //特殊iota赋予常量从0开始的值
		筑基
		结丹
		元婴
		化神
	)
	fmt.Println(练气, 筑基, 结丹, 元婴, 化神)
	fmt.Println(练气 < 结丹, 结丹 > 化神)
}
