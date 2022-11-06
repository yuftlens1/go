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
		case 1: //各个case后的值不能重复
			fallthrough //fallthrough 穿透作用
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
		//case 1, 3, 5, 7, 8, 10, 12:   //这样写也ok。
		//	day = 31
		case 2: //2月默认28天，有时候29天
			if (year%400 == 0) || (year%4 == 0 && year%100 != 0) { // ||两个管道符是或的意思
				day = 29
			} else {
				day = 28
			}
		default: //default有木有都行
			day = 30
		}
		fmt.Printf("%d年%d月有%d天\n", year, month, day)

	} else {
		fmt.Println("月份输入错误")
	}

	//实战2   score得分  grade等级
	//grade := "" //定义一个空的字符串变量
	var grade byte //byte得用单引号引住
	score := 60
	fmt.Println(grade) //byte默认值就是0，这里还是空值,所以会输出0
	fmt.Printf("%T\n", grade)
	switch {
	case score > 90:
		grade = 'A'
	case score >= 80:
		grade = 'B'
	case score >= 70:
		grade = 'C'
	case score >= 60:
		grade = 'D'
	default:
		grade = 'E'
	}
	fmt.Println(grade)             //直接输出byte变量grade,一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。目前还不理解
	fmt.Printf("你的等级是%s\n", grade) //格式化输出字符串变量grade,\n是转义字符--换行

	fmt.Println("你的最终评价是") //根据分数得出等级，根据等级走下面的评价系统得出评价
	switch getAdd(grade) {
	case 'A': //必须加"" 以匹配字符串数据
		fmt.Println("卓越")
	case 'B': //必须加"" 以匹配字符串数据
		fmt.Println("不错")
	case 'C': //必须加"" 以匹配字符串数据
		fmt.Println("可以")
	case 'D': //必须加"" 以匹配字符串数据
		fmt.Println("呦西")
	case 'E': //必须加"" 以匹配字符串数据
		fmt.Println("差点")
	}
}
func getAdd(grade byte) byte { //变量数据类型修改为byte,返回值也修改为byte.这里把byte改成string就不行,byte能逻辑运算应该是
	return grade - 1 //这里减一，但实际效果是加一个等级,我不理解
}

/*
	switch {
	case grade == "A":
		fmt.Println("卓越")
	}
*/
