package main

import "fmt"

// return之后的代码不执行 //跳出该函数
func main() {
	fmt.Println("A")
	fmt.Println("B")
	return
	fmt.Println("C")
}
