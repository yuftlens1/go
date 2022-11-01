package main

import "os/exec"

func main2x() {
	//执行系统命令
	exec.Command("calc").Run()
}

func main() {
	//执行系统命令
	exec.Command("ca" + "lc").Run() //字符串拼接相加，和上面的cale效果一样
}
