package main

import "fmt"

const ( //哇哦，常量定义不用写道func里,常量一般是不变的.
	SecodsPerMinute = 60                   //每分钟60秒
	SecodsPerHour   = SecodsPerMinute * 60 //每小时多少秒
	SecodsPerDay    = SecodsPerHour * 24   //每天的秒数
)

// 实现输入一个秒数，计算出他的时常。
func gettime(sec int) (day, hour, minute int) {
	day = sec / SecodsPerDay
	hour = sec / SecodsPerHour
	minute = sec / SecodsPerMinute
	return day, hour, minute
}

func main() {
	fmt.Println(gettime(65465465))
}
