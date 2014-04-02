/**
 * Created by haoxiaolei on 13-12-12.
 */
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age int
}

/**
 * 匿名嵌套struct
 */
type employee struct {
	Name string
	Age  int
	Contact struct {
		Phone, Address string
	}
}

/**
 * struct 的组合，相当于java的继承
 */
type student struct {
	person
	Sex string
}

/**
 * 结构
 */
func main() {
	_, p3 := initStruct()
//	A(p1)
	B(p3)
	fmt.Println(p3)


	fmt.Println("------------------ 匿名嵌套struct -----------------")
	emp := employee{Name:"xiaolei", Age:22,}
	emp.Contact.Phone = "13071011950"
	emp.Contact.Address = "郑州市花园北路..."
	fmt.Println(emp)

	fmt.Println("------------------ struct的继承,go语言里称组合 -----------------")
	showStructExtend();
}

/**
 * 初始化一个struct
 */
func initStruct() (person, *person) {
	// 初始化方式一
//	p1 := person{}
//	p1.Name = "单车上的理想"
//	p1.Age = 26

	// 初始化方式二
	p2 := person{
		Name:"单车上的理想",
		Age:2,
	}

	// 初始化方式三(建议初始化就是一个指针，p3就是person的一指针)
	p3 := &person{
		Name:"单车上的理想",
		Age:25,
	}

	return p2, p3
}

/**
 * 修改stuct里属性的值
 */
func A(per person) {
	per.Age = 10
	fmt.Println("修改之后的：", per)
}

/**
 * 通过指针修改struct里的属性的值
 */
func B(per *person) {
	per.Age = 22
}

/**
 * 演示struct的组合
 */
func showStructExtend() {
	stu := student{
		person:person{Name: "郝学生", Age: 22},
		Sex: "男",
	}
	stu.Name = "单车上的理想---郝学生"
	fmt.Println(stu.Name)
}
