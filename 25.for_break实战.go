package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand.Intn () 函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (), 并且传入一个变化的值作为参数，如 time.Now().UnixNano() , 就是可以生成时刻变化的值

// break 跳出整个循环体
func main() {
	for true {
		rand.Seed(time.Now().Unix())
		n := rand.Int() % 10 //生成一个10内的随机数!
		fmt.Println(n)
		if n == 9 || n == 2 || n == 3 {
			break // break 跳出整个for循环体
		}
	}

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
}
