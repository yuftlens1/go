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
	return //day, hour, minute
}

/*
// 调用上面写的函数
func main() {
	fmt.Println(gettime(65465465)) // 输出x天  或  x小时  或  x分钟
}
*/

// 升级版，取消 或 关系  //数学得好
func gettimex(seconds int) (day, hour, minute, secondsx int) {
	day = seconds / SecodsPerDay
	seconds -= day * SecodsPerDay // seconds = seconds - (day * SecodsPerDay)
	hour = seconds / SecodsPerHour
	seconds -= SecodsPerHour * hour // seconds = seconds - SecodsPerHour*hour
	minute = seconds / SecodsPerMinute
	secondsx = seconds - minute*SecodsPerMinute
	return
}
func main() {
	fmt.Println(gettimex(119000)) //x秒等于 x天x小时x分钟x秒
}
