package main

import "fmt"

func SUB(a int, b int) (c int) {
	c = a - b
	return c
}

func main() {
	fmt.Printf("%T\n", SUB) //打印出了函数类型，函数本质就是内存地址，类型，调用

	ak47 := SUB              //!!!介个就是函数变量，把函数赋到变量里，一个函数多个名字!!!
	fmt.Printf("%T\n", ak47) //Printf()  //f 格式化输出

	fmt.Println(SUB)
	fmt.Println(ak47)

	fmt.Println(SUB(10, 5))
	fmt.Println(ak47(10, 5))
}

//%T  相应值的类型的Go语法表示
