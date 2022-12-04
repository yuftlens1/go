package main

import "fmt"

/*
一个加法函数
func add(a int, b int) int { //第三个int是定义返回值的数据类型
	return a + b //return返回
}
*/

/*
一个加法和减法函数
func add(a int, b int) (int, int) {
	return a + b, a - b //return返回
}
*/

func add(a int, b int) (x, y int) { //返回值两个x y都是int类型!!!
	return a + b, a - b //return返回
}

func del(a, b int) (c int) { //函数del有a b两个int参数，有一个返回 c int
	c = a - b
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(del(7, 5))
	fmt.Printf("%T\n%T", add, del)
}
