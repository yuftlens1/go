package main

import "fmt"

func main() {
	var a [3]int //定义数组。长度3,分别是 0,1,2位置
	a[1] = 10
	a[2] = 11

	fmt.Println(len(a)) //因为数组a的存储长度是3，编号0 1 2
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[len(a)-1]) // 3-1肯定是2，所以a[len(a)-1] 等于 a[2]
	fmt.Println(a)           //输出数组全部

	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}
