package main

import "fmt"

func main() {
	i := 0
	for {
		if i == 5 { //i等于5就退出了,所以5这层不循环
			break
		}
		fmt.Printf("循环内--%s\n", i)
		i++
	}
}
