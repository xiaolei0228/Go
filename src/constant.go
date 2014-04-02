package main

import (
	"fmt"
)

const a int = 1
const b = 'A'

const (
	c = a
	d = a + 1
	e = a + 6
	f // 如果不赋值，则填充上行的值，即：e的值
	h
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(e)
	fmt.Println(len("sss"))

	fmt.Println(f)
	fmt.Println(h)

	/*****************/
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
