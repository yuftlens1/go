package main

import "fmt"

func main1x() {
	var age int //定义变量，分配内存。
	var tall int
	var heavy int
	age = 19 //为变量赋值
	tall = 170
	heavy = 160
	fmt.Println("老子现在是", age, "岁,身高是:", tall, "体重是:", heavy) //底色是灰色的字符可以当其没有！
}

func mainx() {
	var ( //定义变量，分配内存。
		age   int
		tall  int
		heavy int
	)
	age = 19 //为变量赋值
	tall = 170
	heavy = 160
	fmt.Println("老子现在是", age, "岁,身高是:", tall, "体重是:", heavy) //底色是灰色的字符可以当其没有！
}

func main3x() {
	var ( //定义变量，分配内存。
		age   int = 19
		tall  int = 170
		heavy int = 160
	)
	fmt.Println("老子现在是", age, "岁,身高是:", tall, "体重是:", heavy) //底色是灰色的字符可以当其没有！
}

func main4x() {
	var ( //定义变量，分配内存。
		age, tall, heavy int = 19, 170, 160
	)
	fmt.Println("老子现在是", age, "岁,身高是:", tall, "体重是:", heavy) //底色是灰色的字符可以当其没有！
}

func main() {
	age := 19 //等价于  var _a uint = 10
	tall := 170
	heavy := 160
	fmt.Println("老子现在是", age, "岁,身高是:", tall, "体重是:", heavy) //底色是灰色的字符可以当其没有！

	var a, b int = 10, 20
	fmt.Println(a - b)

	var c = 11.1 //输出float64，证明go会 自动推导数据类型
	d := 13      //等价于 var d int = 13
	fmt.Printf("%T", c, d)
	fmt.Println(d)

	//下面是内存数据库
	aa := 11
	bb := aa
	println(aa, bb) //bb就是aa的备份，下面开始折腾aa，aa折腾坏了还有bb
	aa = 3
	println(aa, bb)

	cc := "woqu"
	fmt.Printf("%T", cc)
}
