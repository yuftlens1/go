package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	// fmt.Println(integer16 + integer32) //会报错,因为不同类型的数据它们不能直接做运算
	fmt.Println(int32(integer16) + integer32) //int32(integer16)就是 把int16类型的数据转为int32

	// Go 的另一种转换方法是使用 strconv 包 //https://blog.csdn.net/u013252047/article/details/104891177
	i, _ := strconv.Atoi("-42") //字符串转整数
	fmt.Printf("i values %v,type is %T\n", i, i)
	s := strconv.Itoa(-42) //整数转字符串
	fmt.Printf("s values %v,type is %T", s, s)
}
