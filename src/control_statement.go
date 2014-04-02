package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------------if...else-----------------------")
	// if...else
	if a := 1; a < 2 {
		fmt.Println("有初始变量...是局部变量")
	}

	fmt.Println("-----------------------for循环的第一种写法----------------------")
	// for
	b := 0
	for {
		b++
		if b > 3 {
			break
		}
		fmt.Println(b)
	}
	fmt.Println("Over")

	fmt.Println("-----------------------for循环的第二种写法----------------------")
	// for 循环的第二种写法
	b1 := 0
	for b1 < 3 {
		b1++
		fmt.Println(b1)
	}
	fmt.Println("Over")

	fmt.Println("-----------------------for循环的第三种写法(强烈推荐)----------------------")
	// for 循环的第三种写法(强烈推荐)
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("Over")

	fmt.Println("-----------------------switch第一种写法----------------------")
	// switch第一种写法
	c := 10
	switch c {
	case 0:
		fmt.Println("c=0")
	case 1:
		fmt.Println("c=1")
	default:
		fmt.Println("none")
	}

	fmt.Println("-----------------------switch第二种写法----------------------")
	// switch第二种写法
	c1 := 10
	switch {
	case c1 == 0:
		fmt.Println("c1=0")
	case c1 >= 1:
		fmt.Println("c1>=1")
		fallthrough // 使用fallthrough可以执行下一行case
	case c1 > 3:
		fmt.Println("c1>3")
	default:
		fmt.Println("none")
	}

	fmt.Println("-----------------------switch第三种写法(在switch后紧跟表达式)----------------------")
	// switch第二种写法(在switch后紧跟表达式)
	switch c2 := 10; {
	case c2 == 0:
		fmt.Println("c2=0")
	case c2 >= 1:
		fmt.Println("c2>=1")
		fallthrough // 使用fallthrough可以执行下一行case
	case c2 > 3:
		fmt.Println("c2>3")
	default:
		fmt.Println("none")
	}

	/*********************goto，break，continue 跳转语句*********************/
	// break
	fmt.Println("-----------------------for循环使用break LABEL跳出----------------------")
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("成功使用break label跳出循环")

	// continue
	fmt.Println("-----------------------for循环使用continue LABEL跳出，用在死循环中，死循环会立即停止----------------------")
LABEL2:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL2
			fmt.Println(i) // 永远执行不到
		}
		fmt.Println(i) // 永远执行不到
	}
	fmt.Println("成功使用continue LABEL跳出循环")

	// goto
	fmt.Println("-----------------------for循环使用goto LABEL跳出,LABEL必须放循环后，否则是死循环----------------------")
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL3
			}
		}
	}
LABEL3:
	fmt.Println("成功使用goto label跳出循环")
}
