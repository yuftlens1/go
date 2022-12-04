package main

import (
	"fmt"
	"os/exec"
)

// 函数：功能代码块。函数解决的问题：代码重用。
func mainh() {
	fmt.Println("要启动记事本了")
	fmt.Println("notepad")
	exec.Command("notepad").Run()

	fmt.Println("要启动计算器了")
	fmt.Println("calc")
	exec.Command("calc").Run()
}

// 使用函数简化上面的代码块
func gocmd(cmd string) { //定义函数gocmd,定义字符串变量cmd!
	fmt.Printf("要启动%s了\n", cmd)
	fmt.Println(cmd)
	exec.Command(cmd).Run()
}
func main() {
	gocmd("notepad") //调用函数gocmd
}

/*
func 函数名(){
	函数体
}
*/
