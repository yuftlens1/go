package main

import "fmt"

//实现效果，为数据加序号，  这个叫     ！切片迭代！

func main() {
	num1 := []int{1, 2, 3} //定义列表
	fmt.Println(num1)
	for k, v := range num1 { //v := range num1  看成一体
		fmt.Printf("key: %v , value: %v  \n", k, v)
		fmt.Printf("key: %d , value: %d  \n", k, v)
	}

	mystr := "hello,中国"           //字符型数据变量
	for value, i := range mystr { //i := range mystr 看成一体
		fmt.Printf("%c%d\n", i, value)
	}

	for i := 0; i < len(mystr); i++ {
		fmt.Printf("%c\n", mystr[i]) // mystr[i] // [] 应该是输出数据的内存下标.字符串/数据切片 。切片不能输出字符串不知道为啥，英文就可以
	}
	fmt.Println(mystr[0])

	var str string
	str = "qwertyuiop"
	fmt.Printf("%T\n", str)
	fmt.Println(str[0])
	fmt.Printf("%c\n", str[0])
	fmt.Print(str)
}

/*
%v 相应值的默认格式。
%d 一个十进制数值(基数为10)
%c 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符
%s 字符串。输出字符串中的字符直至字符串中的空字符（字符串以'\0‘结尾，这个'\0'即空字符）
*/
