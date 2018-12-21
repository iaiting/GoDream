package main

import (
	"fmt"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////
//
// Author
// Generate Data	:
// 第 2 章 顺序编程
//
////////////////////////////////////////////////////////////////////////////////
func chap02_t00() {

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
func chap02_t01() {
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

// 2.2 常量
// 2.2.2 常量定义
func chap02_t02() {
	simple_log("Enter:")

	const zero_0 = 0.0
	const zero_1 = 0

	const zero_2 = iota

	const v1 = iota * 11

	fmt.Println("zero_0 类型:", reflect.TypeOf(zero_0))
	fmt.Println("zero_1 类型:", reflect.TypeOf(zero_1))
	fmt.Println("zero_2 类型:", reflect.TypeOf(zero_2))
	fmt.Println("v1 类型:", reflect.TypeOf(v1))

	simple_log("Leave:\n")
}

// 星期枚举
func chap02_t03() {
	simple_log("Enter:")

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println("Sunday =", Sunday)

	fmt.Println("Monday =", Monday)

	fmt.Println("Tuesday =", Tuesday)

	fmt.Println("Wednesday =", Wednesday)

	simple_log("Leave:\n")

}

// 2.3 类型
// 2.3.1 布尔类型
func chap02_t04() {
	simple_log("Enter:")

	var b_v1 bool = true
	var b_v2 bool = false
	fmt.Println("b_v1 =", b_v1)
	fmt.Println("b_v2 =", b_v2)

	simple_log("Leave:\n")

}

// 2.3.2 整型
func chap02_t05() {
	simple_log("Enter:")

	var i_v1 int8 = 127
	fmt.Println("i_v1 =", i_v1)

	var ui_v2 uint8 = 0xfe
	fmt.Println("i_v1 =", ui_v2)

	simple_log("Leave:\n")
}

// 2.3.3 浮点型
func chap02_t06() {
	simple_log("Enter:")

	var fval1 float32
	fval1 = 12
	fmt.Println("fval =", fval1)

	fval2 := 13.1
	fval2 = float64(fval1)

	fmt.Println("fval2 =", fval2)

	simple_log("Leave:\n")
}

// 2.3.4 复数型
func chap02_t07() {
}

////////////////////////////////////////////////////////////////////////////////
func chap02_main() {
	chap02_t01()
	chap02_t02()
	chap02_t03()
	chap02_t04()
	chap02_t05()
	chap02_t06()
}
