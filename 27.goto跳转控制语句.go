package main

//goto:  执行指定的代码!!!

import "fmt"

func main() {
	i := 0
AK:
	if i < 5 { //这本来不是循环，就是因为goto+AK成了循环    //如果 i < 5才执行下面的！
		i++
		fmt.Println(i)
		goto AK
	}

	goto HAHA
	fmt.Println("A")
	fmt.Println("B")
HAHA:
	fmt.Println("C")
	fmt.Println("D")
	fmt.Println("E")
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
