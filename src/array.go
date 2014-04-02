package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [20]int{19: 1}
	fmt.Println(b)

	var p *[20]int = &b
	fmt.Println(*p)

	x, y := 1, 2
	c := [...]*int{&x, &y}
	fmt.Println(c)

	// 给数组赋值
	arr1 := [10]int{}
	arr1[1] = 88
	fmt.Println(arr1)
	arr2 := new([10]int)
	arr2[1] = 88
	fmt.Println(arr2)

	// 多维数据
	aa := [2][3]int{{1, 1, 1}, {2, 2, 2}}
	fmt.Println(aa)
	bb := [2][3]int{{2: 88}, {2, 2, 2}}
	fmt.Println(bb)
}
