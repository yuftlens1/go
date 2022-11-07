package main

import (
	"fmt"
)

// nil空值
func main() {
	var x interface{} //interface代表万能接口，这里是接收任何类型。这里替换了传统的 var x int 定义方法

	y := 3 > 2 //数据源头
	x = y
	fmt.Println(x)

	switch i := x.(type) {
	case int:
		fmt.Println("int类型")
	case float32:
		fmt.Println("float32类型")
	case float64:
		fmt.Println("float64类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("布尔类型")
		fmt.Println(i)
	case nil:
		fmt.Printf("%T", i)
	default:
		fmt.Println("未匹配类型")
	}

}
