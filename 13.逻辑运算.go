package main

import "fmt"

// && 两边的条件全符合为true。  || (或) 满足一个条件即可

func main() {
	var (
		期望颜值 int = 100
		期望财富 int = 100

		我的颜值 int = 99
		我的财富 int = 100
	)
	if 我的颜值 >= 期望颜值 && 我的财富 >= 期望财富 {
		println("嫁了")
	} else {
		println("滚远")
	}

	var (
		可以支付宝 bool = !!!false //多个同级逻辑运算：从右往左
		可以微信  bool = !true    //!是取反的意思，这里会把true变成false!!!
	)
	fmt.Println(可以支付宝)
	if 可以微信 || 可以支付宝 {
		println("可以结账")
	} else {
		println("不好结账")
	}
}
