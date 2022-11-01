package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(&a)    //输出a的内存地址
	fmt.Println(*(&a)) //根据内存地址取出内容

	var p *int = &a    //*int指针
	fmt.Println(p)     //输出内存地址
	fmt.Println(*(&p)) //还是输出内存地址

	*p = 2 //修改同一片内存
	fmt.Println(a, *p, *&a)
}
