package main

import "fmt"

func go1() {
	fmt.Println("go1 up")
	fmt.Println("go1 down")
	//return 1 //return 对应返回值那里！
}

func go2() int {
	fmt.Println("go2 up")
	fmt.Println("go2 down")
	return 2 //return 对应返回值那里！
}

func go3() int {
	fmt.Println("go3 up")
	go2() //函数里调用函数，和shell一样
	fmt.Println("go3 down")
	return 3 //return 对应返回值那里！
}

func main() { //main函数是执行体
	go1() //直接执行go1

	//fmt.Println(go2)
	fmt.Println(go2()) //上面go后不加()的话就输出go2的内存地址了，加了()就输出go2 的返回值了。

	go3() //直接执行go3   //直接执行不会输出return，因为没有print打印输出，实际返回值是有的！！！

	//下面是嵌套调用

}

func DIV(a, b int) int {
	return a / b
}
