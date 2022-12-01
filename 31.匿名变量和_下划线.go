package main

import "fmt"

// go里面定义变量后不使用会报错!
// _ 下划线是占位符/是匿名变量，不可以使用，功能是抵消无用数据！匿名变量不会占内存也不会分配内存

func GetData() (int, int) { //这样的函数会吐数据出来
	return 100, 200
}

func main() {
	a, b, c, d, _ := 1, 2, 3, 5, 6 //key与value在数量上必须一致对其 // _ 是匿名变量，不可以使用，功能是抵消无用数据！
	// fmt.Println(a, b, c, d, _)  //使用是匿名变量会报错!
	fmt.Println(a, b, c, d)

	//在 go里，有的函数返回多个变量,但只用一个怎么办，匿名变量这时候就派上用场了
	haha, _ := GetData()
	_, hehe := GetData()
	fmt.Println(haha, hehe)

	w, q := GetData()
	fmt.Println(w, q)
}
