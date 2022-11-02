package main

import "fmt"

// 判断x年x月有多少天
func main() {
	var (
		year  int
		month int
		day   int
	)
	fmt.Println("请输入年份")
	fmt.Scanf("%d\n", &year) //输入在变量前一定要加 & !!!
	fmt.Println("请输入月份")
	fmt.Scanf("%d\n", &month)

	if month >= 1 && month <= 12 {
		switch month {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			day = 31
		case 2: //2月默认28天，有时候29天
			if (year%400 == 0) || (year%4 == 0 && year%100 != 0) { // ||两个管道符是或的意思
				day = 29
			} else {
				day = 28
			}
		default:
			day = 30
		}
		fmt.Printf("%d年%d月有%d天", year, month, day)

	} else {
		fmt.Println("月份输入错误")
	}
}
