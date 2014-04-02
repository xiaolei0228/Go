package main

import (
	"fmt"
	"sort"
)

/**
 * map就是K-V形式来存储数据的，key必须是支持==或!=比较运算的类型，不可是函数、map或slice
 * map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
 * map使用make()创建，支持:=这种简写方式
 *
 * make([keyType]valueType, cap),cap表示容量，可省略，超出容量时会自动扩容，但尽量合理初始化容量
 * 使用len()获取元素个数
 *
 * 键值对不存在时自动添加
 * 使用delete()删除某键值对
 * 使用for range对map和slice进行迭代操作
 */
func main() {
	makeMap()

	deleteMap()

	complexMap()

	iterationMap()

	sortMap()
}

/**
 * 创建map
 */
func makeMap() {
	// var m map[int]string = make([int]string)  // 方式一
	m := make(map[int]string) // 方式二
	m[1] = "OK"
	fmt.Println(m)
}

/**
 * 删除map里的元素
 */
func deleteMap() {
	m := make(map[int]string)
	m[1] = "OK"
	delete(m, 1)
	fmt.Println(m)

	// 直接赋值map
	m2 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m2)
}

/**
 * map的嵌套,复杂map的处理
 */
func complexMap() {
	m := make(map[int]map[int]string)

	a, ok := m[2][1]
	if !ok { // 如果不存在就make出来
		m[2] = make(map[int]string)
		m[2][1] = "GOOD,动态创建出来map了,是否创建成功，用ok来判断"
		a, ok = m[2][1]
	}

	fmt.Println(a, ok)
}

/**
 * 迭代map/slice
 */
func iterationMap() {
	sm := make([]map[int]string, 5) // 声明一个元素类型为map的slice
	for i := range sm {             // 当不需要索引的时候用_来代替
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}

/**
 * map排序
 */
func sortMap() {
	m := map[int]string{3: "c", 4: "d", 1: "a", 5: "e", 2: "b"}

	s := make([]int, len(m))
	count := 0
	for k, _ := range m {
		s[count] = k
		count++
	}
	sort.Ints(s)

	// sort map
	for i := 0; i < len(m); i++ {
		fmt.Println(m[s[i]])
	}
	fmt.Println(s)
}
