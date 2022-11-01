package main

import "fmt"

func main() {
	var dole int
	fmt.Println("请输入你的美元数")
	fmt.Scanf("%d", &dole)

	switch { //如果需要判断大小就不能在这里写变量
	case dole < 999:
		fmt.Println("穷光蛋蛋")
		fallthrough //加了这个后,命中这个case的话会执行这个case和下一个case
	case dole < 9999 && dole >= 999:
		fmt.Println("穷光蛋")
	case dole < 99999 && dole >= 9999:
		fmt.Println("有钱人")
	default:
		fmt.Println("switch丫")
	}
}
