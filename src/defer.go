/**
 * Created by haoxiaolei on 13-12-12.
 */
package main

import (
	"fmt"
)

/**
 * defer：
 *		1、执行方式类似于其它语言的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
 * 		2、即使函数发生【严重错误】也会执行
 *		3、支持匿名函数的调用
 *		4、常用于资源清理、文件关闭，解锁以及记录时间等操作
 *		5、通过与匿名函数配合可以return之后修改函数计算结果
 *		6、如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已获得了拷贝，
 *		   否则则是引用某个变量的地址
 *
 * panic/recover：
 *		Go语言中没有异常机制，但有panic/recover模式来处理错误
 *		Panic可以在任何地方引发，但recover只有在defer调用的函数有效
 */
func main() {
	//	testDefer()

	A()
	B()
	C()
}

func testDefer() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func A() {
	fmt.Println("func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recocver in B")
		}
	}()
	panic("panic in B")
}

func C() {
	fmt.Println("func C")
}

