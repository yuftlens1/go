package main

import "fmt"

//标识符用来命名变量、类型.

func main() {
	var _a uint = 10              //定义一个变量a=数值10。 a就是标识符.  标识符必须以字母或者下划线开头
	fmt.Println("老子现在是", _a, "岁") //底色是灰色的字符可以当其没有！
	_a = _a + 5
	fmt.Println("老子5年后", _a, "岁")

	const name string = "老子" //定义一个不能改变的常量。变量 _a 一直在变。
	_a = _a + 10
	fmt.Println("10年后", name, _a, "岁")
}
