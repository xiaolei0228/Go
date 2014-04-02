package main

import (
	"fmt"
)

/**
 * 位运算符

  6: 0110
 11: 1011
-----------
 &   0010 = 2	(与)
 |   1111 = 15	(或)
 ^   1101 = 13	(在对比过程当中，两个只有一个为1的时候才成立)
 &^  0100 = 4	(在两个数比较的时候，当第二个数的某个位上值为1，则输出第一个数的对应位上的值改为0)
*/
func main() {
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	a := 0
	if a > 0 && (10/a) > 1 { // 说明了&&运算符的使用，当&&前面的不成立，后面的表达式就不会执行了
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}
}
