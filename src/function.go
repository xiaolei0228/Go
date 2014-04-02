/**
 * Created by haoxiaolei on 13-12-11.
 */
package main

import (
	"fmt"
)

/**
 * 函数
 *		不支持------ 嵌套、重载、默认参数
 *		支持----- 无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
 *
 * 定义函数使用关键字 func，且左大括号不能另起一行
 * 函数也可以做为一种类型使用
 *
 * 示例： func 函数名(参数名1  参数类型，参数名1  参数类型...) (返回值1的类型，返回值2的类型...) {...函数体...}
 * 		如果返回值只有一个，返回值类型不用带括号
 * 		参数
 *
 */
func main() {
	//	C(1,2,3,4,5,6)

	// 调用闭包函数
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(5))
}

/**
 * A函数要a和b两个参数，返回值是
 */
func A(a string, b int) {
}

/**
 * 命名返回值函数，返回值类型如果都是同一种类型，可以用这种方式
 */
func B() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

/**
 * 不定长参数，这样的参数只参放到参数的最后，相当于一个slice了，
 * 也相当于java中的不定长参数，属于一个参数容器
 */
func C(a ...int) {
	fmt.Println(a)
}

/**
 * 闭包函数
 */
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
