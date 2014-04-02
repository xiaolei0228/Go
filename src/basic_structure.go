// 当前程序的包名
package main

// 导出其它的包
import (
	"fmt"
)

// import std "fmt"   /** 给包定义别名std **/

// 常量的定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct {
}

// 接口的声明
type golang interface {
}

// 由main函数作为程序入口点启动
func main() {
	fmt.Println("Hello world！你好，世界！")
}
