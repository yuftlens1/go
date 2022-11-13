package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for n := 0; n < 3; n++ { //控制多少列，因为 Print 不换行
			fmt.Print("*")
		}
		fmt.Println("*") //控制多少行.	因为 ln 换行
	}

	for ii := 0; ii <= 9; ii++ {
		/*
			for j := 0; j <= 9; j++ { //这样ii每次循环都会把j这层循环跑满,但是下面加了if判断就不会跑满了！
				if j >= ii {
					fmt.Printf("%d乘%d等于%d   ", ii, j, ii*j)
				}

		*/
		for j := 0; j <= ii; j++ { //等价于上面的 里面层的for+if !!!
			fmt.Printf("%d乘%d等于%d   ", ii, j, ii*j)
		}
		fmt.Println() //只为换行
	}
}
