package main

import (
	"fmt"
)

func main() {
	a := [...]int{12, 2, 4, 1, 9}
	fmt.Println(a)

	count := len(a)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
