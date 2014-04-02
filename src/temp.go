package main

import (
	"fmt"
	"math"
	"strconv"
)

type (
	byte int8
	rune int32
	文本   string // 建议不使用
)

func main() {
	var a [1]byte
	fmt.Println(a)

	fmt.Println(math.MaxInt32)

	var b 文本
	b = "中文类型名"
	fmt.Println(b)

	// 声明变量 ---方式1(可以省略变量类型，由系统推断)
	var c = 1.1
	fmt.Println(c)
	// 声明变量 ---方式2(最简写法)
	d := 2.2
	fmt.Println(d)

	/****************************/
	// var a1, a2, a3, a4 int = 1, 2, 3, 4
	// 或
	a1, a2, a3, a4 := 1, 2, 3, 4
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	// 相互兼容的两种类型之间的转换
	var cc float32 = 1.1
	ccc := int(cc)
	fmt.Println(ccc)

	// int类型转string类型会转成字母
	var aa int = 65
	bb := string(aa)
	fmt.Println(bb)
	fmt.Println(strconv.Itoa(aa))

}
