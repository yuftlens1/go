package main

import "fmt"

func add1(a, b, c, d int) int { //定义形参a b c d
	return a + b + c + d
}
func add2(a, b, c, d int) (e int) {
	e = a + b + c + d
	return e
}
func add3(a, b, c, d int) (e int) {
	a = 100 //函数里的变量替换了实参！上面已经定义了a是int的数据类型,函数内的变量值优先级较高
	e = a + b + c + d
	return
}

func main() {
	fmt.Println(add1(1, 1, 1, 1)) //传入实参
	fmt.Println(add2(1, 1, 1, 1))
	fmt.Println(add3(1, 1, 1, 1))

	x, y, z, m := 2, 3, 4, 5
	fmt.Println(add2(x, y, z, m))
}
