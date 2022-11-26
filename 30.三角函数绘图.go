package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	const size = 300                                   //定义常量 size;图像大小
	pic := image.NewGray(image.Rect(0, 0, size, size)) //根据给定大小创建灰度图。Gray灰色;灰度

	//遍历每个像素，为其填充颜色 //生成画板！！！
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//生成x坐标 //画图  //在画板里画线！！！
	for xx := 0; xx < size; xx++ {
		s := float64(xx) * 2 * math.Pi / size
		yy := size/2 - math.Sin(s)*size/2
		pic.SetGray(xx, int(yy), color.Gray{0}) //黑色sin线
	}

	//画好的图生成文件
	file, _ := os.Create("sin.png")
	png.Encode(file, pic)
	file.Close()
}
