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
	//return a + b, a - b //return返回   //这样写也是可以的

	x = a + b
	return x, a - b //return返回
}
func del(a, b int) (c int) { //函数del有a b两个int参数，有一个返回 c int
	c = a - b
	return c
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(del(7, 5))
	fmt.Printf("%T\n%T\n", add, del)

	num1, num2 := add(2, 3) //把add函数输出的x y 接收到变量 num1,num2 里面
	fmt.Println(num1, num2)
}
