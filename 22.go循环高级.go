package main

import "fmt"

func main() {
	ii := 0
	for {
		if ii == 5 { //i等于5就退出了,所以5这层不循环
			break
		}
		fmt.Printf("循环内--%d\n", ii)
		ii++
	}

	mystr := "hello,中国" //字符型数据变量
	for i, value := range mystr {
		fmt.Printf("%d%c\n", i, value)
	}

	var str string
	str = "hello world"
	fmt.Println(str[0:4]) //字符串切片, 4之前
	fmt.Println(str[:4])  //和上面效果一样
	fmt.Println(len(str))
}

/*
%d 一个十进制数值(基数为10)
%c 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符
%s 字符串。输出字符串中的字符直至字符串中的空字符（字符串以'\0‘结尾，这个'\0'即空字符）
*/
