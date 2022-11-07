package main

import (
	"fmt"
	"os/exec"
)

func main() {
	/*
		for i := 0; i < 10; i++ { //++的效果就是+1,自增 //在go里定义变量  var qaq int = 99 等价于 qaq := 99
			fmt.Println("haha")
			fmt.Printf("%T", i)
		}  */
	for i := 5; i < 12; i += 2 { //i+=2每次加2 (i=i+2)    //在go里定义变量  var qaq int = 99 等价于 qaq := 99
		fmt.Println("haha")
	}

	for true { //   和 for {  的效果一样，死循环
		exec.Command("notepad").Run()
	}
}
