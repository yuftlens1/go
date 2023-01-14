package main

import "fmt"

func shouDataType(list ...interface{}) {
	for _, arg := range list { // _ 占位符，这里替换的是数组的位置序号!把值赋给 arg
		switch arg.(type) {
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
		case nil:
			fmt.Printf("%T", arg)
		case byte, rune:
			fmt.Println("字符类型")

		default:
			fmt.Println("未匹配类型")
		}
	}
}

func main() {
	shouDataType("ok", 1, 1.212, "X", false)

	/*
		a := "ok"
		fmt.Printf("%T", a)
	*/
}
