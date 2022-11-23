package main

import "fmt"

func main() {
	i := 0
AK:
	if i < 50 {
		i++
		fmt.Println(i)
		goto AK
	}
}

/*
   AK:
   	for i := 0; i < 5; i++ {
   		for j := 0; j < 9; j++ {
   			if j > 5 {
   				break AK //不加只跳出了内层j的for循环 //加了AK 就从前面AK i 外层停止循环
   			}
   			fmt.Printf("*")
   		}
   		fmt.Println("*")
   	}
   	fmt.Println()
*/
