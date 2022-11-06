package main

import "fmt"

func main1() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)
}

func main() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	//fmt.Println(integer16 + integer32) //不同类型的数据它们不能直接做运算，需要做数据类型转换才能实现运算
	fmt.Println(int32(integer16) + integer32) //数据类型转换，把int16转为int32

	var test byte
	test = 129
	fmt.Println(test)
}
