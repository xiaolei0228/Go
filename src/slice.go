package main

import (
	"fmt"
)

/**
 * slice 是指向底层数据的指针(只有在底层数据的内存地址不发生改变的时候成立)
 * 其容量cap从开始位置一直计算到底层数据结束
 *
 */
func main() {
	// slice的声明
	var sli []int
	fmt.Println(sli)

	// 从数组中取数据填充到slice
	a := [10]int{1, 3, 6, 2, 32, 6, 88, 9, 22, 132}
	sl := a[0:5] // 取数组a中的前5个索引元素
	fmt.Println(sl)

	sl1 := a[:3] // 从开头索引元素取到索引3
	fmt.Println(sl1)

	sl2 := a[6:] // 从第6个索引位置元素一直取到数组末尾
	fmt.Println(sl2)

	makeSlice()

	useSlice()

	appendElemntsIntoSlice()

	copySlice()
}

/**
 * 用make来声明slice
 */
func makeSlice() {
	/**
	 * make(类型，长度，在内存中的连续容量大小)
	 * 如果容量不指定，默认就是长度的大小
	 */
	sl := make([]int, 3, 10)
	fmt.Println(len(sl), cap(sl))
}

/**
 * 使用slice
 */
func useSlice() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5] // 数组截取是左闭右开区间范围
	fmt.Println(len(sa), cap(sa))
	fmt.Println(string(sa))
}

/**
 * 追加元素到slice
 */
func appendElemntsIntoSlice() {
	sl := make([]int, 3, 6) // 声明一个长度为3，容量为6的slice
	fmt.Println("%p", sl)
	sl = append(sl, 1, 2, 3)
	fmt.Println("%v %p", sl, sl)
}

/**
 * 复制slice
 * copy多少个数据，取决于copy对象的容量大小
 * 数组copy时位置会一一对应
 */
func copySlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	s3 := make([]int, 6, 6)
	copy(s2, s1)
	fmt.Println(s2)

	copy(s3, s1)
	fmt.Println(s3)
}
