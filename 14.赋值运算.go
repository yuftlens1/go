package main

import "fmt"

func main() {
	var a int = 3
	a = 12 //重新赋值
	fmt.Println(a)

	a += 2 //等价于 a = a+2
	fmt.Println(a)

	a -= 2 //等价于 a = a-2
	fmt.Println(a)

	a *= 2 //等价于 a = a*2
	fmt.Println(a)

	a /= 2 //等价于 a = a/2
	fmt.Println(a)

	a %= 5 //等价于 a = a%5  //求余
	fmt.Println(a)
}
