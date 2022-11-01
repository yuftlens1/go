package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd string   //为变量cmd开辟内存
	fmt.Scanln(&cmd) //键盘输入,必须有&    //把键盘输入的内容赋到字符串变量cmd上
	fmt.Println(cmd)

	//if单分支
	if cmd == "calc" {
		fmt.Println("是计算器丫,if单分支体验")
		exec.Command(cmd).Run() //执行系统命令
	}

	//if双分支
	var money int
	fmt.Println("请输入你的金钱数")
	fmt.Scanln(&money)

	if money >= 99999 {
		fmt.Println("有钱人")
	} else {
		fmt.Println("穷光蛋")
	}

	//if多分支!   超级富豪,大富豪,有钱人,穷光蛋
	var dole int
	fmt.Println("请输入你的美元数")
	fmt.Scanln(&dole)

	if dole < 999 {
		fmt.Println("穷光蛋")
	} else if dole >= 999 && dole < 9999 {
		fmt.Println("有钱人")
	} else if dole >= 9999 && dole < 99999 {
		fmt.Println("大富豪")
	} else {
		fmt.Println("超级富豪")
	}
}
