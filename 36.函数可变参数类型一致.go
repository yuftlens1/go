package main

import "fmt"

//#  ...int 可变形参要放到 形参括号内的最后面，例如: func kb(a int,b int,args ...int)

func kb(args ...int) int { //...int 可变参数，类型一致 //args就是一个数组
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i] //sum = sum + args[i]  //每一个元素
	}
	return sum
}

/*
func main() {
	fmt.Println(kb(1, 2, 3)) //函数ke 可以无限输入实参，因为形参那边是个数组在接收实参。
}
*/
// i,data := range args   //把数组 args 的位置序号赋值给i, 数据赋值给data
func mul(args ...int) int { //...int 可变参数，类型一致 //args就是一个数组
	sum := 1 //得是1开始,不然0*任何数都是0了
	/*
		for i, data := range args {
			sum *= data //每一个元素  //sum = sum * data
			fmt.Println("->", "位置:", i, "数据:", data)
		}
	*/
	for _, data := range args { //数组切片的位置序列号不参与计算，所以可以用_替换掉(可以节省内存)!!!
		sum *= data //每一个元素  //sum = sum * data
		fmt.Println("->", "数据:", data)
	}
	return sum
}
func main() {
	fmt.Println("数据相乘", mul(1, 2, 5, 3))
}
