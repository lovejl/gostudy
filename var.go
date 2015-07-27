package main

import (
	"fmt"
)

var (
	v1  int       = 10
	v2            = 20
	v3            = "语言"
	v4  [3]int    //整型数组
	v11 [4][2]int //3行5列的二维整型数组
	v5  []string
	v6  *int
	v7  map[string]int
	v8  func(x, y int) int
	v9  bool
)

var v10 struct {
	name string
	age  int
}

func main() {
	v1, v2 = v2, v1
	fmt.Println(v1, v2)
	//v9 = 1 这样写是不对的
	v9 = false
	v9 = (1 == 1)
	fmt.Println(v9)
	v1 = 5 % 3
	fmt.Println(v1)
	fmt.Println(v3, v3[1], len(v3), v3[0:6])
	for i := 0; i < 6; i++ {
		fmt.Println(i, v3[i])
	}
	for i, ch := range v3 { //每个字符的类型是rune(早期的Go语言用int类型表示Unicode字符)而不是byte。
		fmt.Println(i, ch)
	}
	fmt.Println(len(v11))
	array := [3][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	for i, v := range array {
		fmt.Println(i, v)
	}
	for i, v := range v11 {
		fmt.Println(i, v)
	}
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice), "//空间不够随机分配了空间")
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[2:5]
	fmt.Println(newSlice)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
}
