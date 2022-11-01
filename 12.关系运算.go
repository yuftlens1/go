package main

import "fmt"

func main() {
	const (
		练气 = iota //特殊iota赋予常量从0开始的值
		筑基
		结丹
		元婴
		化神
	)
	fmt.Println(练气, 筑基, 结丹, 元婴, 化神)
	fmt.Println(练气 < 结丹, 结丹 > 化神) //关系运算。成立是true，反之false
	fmt.Println(练气 == 结丹)
	fmt.Println(结丹 == 结丹)
	fmt.Println(练气 != 结丹) //!=  不等于的意思
	fmt.Println(筑基 >= 化神)
	fmt.Println(化神 <= 化神)
}
