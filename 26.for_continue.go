package main

import "fmt"

func main() {
	var y int = 0
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i) //打印10之内的奇数
		y = y + i      //计算10之内的奇数总和，算是把每次循环的结果都攒起来
	}
	fmt.Println(y) //输出10之内的奇数总和

	//实战：正负数统计
	var zhengshu int32 = 0
	var fushu int32 = 0
	for {
		/*
			var input int             //只是定义int类型的input变量
			fmt.Scanf("%d\n", &input) //Scanf在win必须加\n!!!  Windows的默认换行符号是\r\n
			if input > 0 {
				zhengshu++
			} else if input < 0 {
				fushu++
			} else {
				break
			}
			fmt.Println(zhengshu, fushu)
		*/

		fmt.Println(zhengshu, fushu) //要想把输出放上面就在每一个判断后加continue。捞到就拿上，不往下走！
		var input int                //只是定义int类型的input变量
		fmt.Scanf("%d\n", &input)    //Scanf在win必须加\n!!!  Windows的默认换行符号是\r\n
		if input > 0 {
			zhengshu++
			continue
		} else if input < 0 {
			fushu++
			continue
		} else {
			break
		}
	}
}

//continue 跳出这一次循环，这一次循环不执行
