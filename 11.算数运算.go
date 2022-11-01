package main

import "fmt"

func main() {
	fmt.Println(10+2, 10-2, 10*2, 10/2) //go里的加减乘除
	fmt.Println(10 % 3)                 //能否被整除，能否被平均

	var timeya = 100
	timeya++ //++的效果就是+1,自增
	timeya++
	println(timeya)
	timeya-- //++的效果就是-1,自减
	println(timeya)
}
