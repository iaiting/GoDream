package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
//
// Author
// Generate Data	:
// 第 2 章 顺序编程
//
////////////////////////////////////////////////////////////////////////////////
func chap02_t01() {

	simple_log("Enter:")

	const c1 float32 = 3.14

	var v1 int = 12

	const eof = 0

	fmt.Println(fmt.Sprintf("%T", v1))
	fmt.Println(fmt.Sprintf("%T", c1))
	fmt.Println(fmt.Sprintf("%T", eof))

	simple_log("Leave:\n")
}

// 2.1.1 变量声明
func chap02_t0101() {
	simple_log("Enter:")

	// 整型变量声明
	var v1 int

	// 字符串变量声明
	var v2 string

	// 整型数组变量声明
	var v3 [10]int

	// 整型数组切片变量声明
	var v4 []int

	// 结构体变量声明
	var v5 struct {
		v int
	}

	// 指针变量声明
	var v6 *int

	// map变量声明
	var v7 map[string]int

	// 函数变量声明
	// var v8 func(int) int

	// 多个变量一起声明
	var (
		v9  int
		v10 string
	)

	fmt.Println("整型变量 v1:", v1)
	fmt.Println("字符串变量 v2:", v2)
	fmt.Println("整型数组变量 v3:", v3)
	fmt.Println("整型数组切片变量 v4:", v4)
	fmt.Println("结构体变量 v5:", v5)
	fmt.Println("指针变量 v6:", v6)
	fmt.Println("map变量 v7:", v7)
	// fmt.Println("map变量 v8:", v8)
	fmt.Println("map变量 v9:", v9)
	fmt.Println("map变量 v10:", v10)

	simple_log("Leave:\n")
}

////////////////////////////////////////////////////////////////////////////////
func chap02_main() {
	// chap02_t01() //
	chap02_t0101()
}
