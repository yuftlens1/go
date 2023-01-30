package main

import "fmt"

func changenum(num int) {
	num = 1000 //函数内的变量优先级较高，所以下行打印的是1000!
	fmt.Println(num, "changenum", &num)
}

func main() {
	num := 100
	changenum(num)
	fmt.Println(num, "main", &num)
}
