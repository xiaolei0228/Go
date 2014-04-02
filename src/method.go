/**
 * Created by haoxiaolei on 13-12-13.
 */
package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type TZ int

func main() {
	a := A{}
	a.Print()

	b := B{}
	b.Print()

	var tz TZ
	tz.Print()
}

/**
 * struct A 的方法
 */
func (A) Print() {
	fmt.Println("this is A.Print()...")
}

func (B) Print() {
	fmt.Println("this is B.Print()...")
}

/**
 * 为底层类型的type 添加方法
 */
func (a *TZ) Print() {
	fmt.Println("TZ...", a)
}
