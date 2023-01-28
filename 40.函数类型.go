package main

import "fmt"

func SUB(a int, b int) (c int) {
	c = a + b
	return c
}

// 下面是函数类型  type定义一个类型
type MyFuncType func(int, int) int //!!!定义函数类型

func main() {
	var fi MyFuncType = SUB
	fmt.Println(fi(1, 2))
	fmt.Printf("%T%V\n", fi, fi)
	fmt.Printf("%T%V\n", SUB, SUB)
}

//函数类型和函数变量多少有点脱裤放屁
