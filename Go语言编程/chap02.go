package main

import (
	"fmt"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////
//
// Author			: iaiting
// Generate Data	: 2018-12-21
// Descriptin		: 《Go语言编程》第 2 章 顺序编程
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

	// 匿名结构体变量声明
	var v5 struct {
		i int
	}

	// 整型指针变量声明
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
	fmt.Println("匿名结构体变量 v5:", v5)
	fmt.Println("指针变量 v6:", v6)
	fmt.Println("map变量 v7:", v7)
	// fmt.Println("map变量 v8:", v8)
	fmt.Println("多变量中的int变量 v9:", v9)
	fmt.Println("多变量中的string变量 v10:", v10)

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

	var b_v3 bool = True
	var b_v4 bool = Ture
	var b_v5 bool = ture

	var b_v6 bool = False
	var b_v7 bool = Flase
	var b_v8 bool = flase

	fmt.Println("b_v1 =", b_v1)
	fmt.Println("b_v2 =", b_v2)

	fmt.Println("b_v3 =", b_v3)
	fmt.Println("b_v4 =", b_v4)
	fmt.Println("b_v5 =", b_v5)

	fmt.Println("b_v6 =", b_v6)
	fmt.Println("b_v7 =", b_v7)
	fmt.Println("b_v8 =", b_v8)

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

	var fv1 float32
	fv1 = 12.0
	fmt.Println("fv1 =", fv1)

	fv2 := 13.1
	fv2 = float64(fv1)

	fmt.Println("fval2 =", fv2, reflect.TypeOf(fv2))

	simple_log("Leave:\n")
}

// 2.3.4 复数型
func chap02_t07() {
	simple_log("Enter:")

	var v1 complex64
	v1 = 1 + 2i
	fmt.Println("v1 =", v1, reflect.TypeOf(v1))

	v2 := 2 + 3i
	fmt.Println("v2 =", v1, reflect.TypeOf(v2))

	v3 := complex(4, 5)
	fmt.Println("v3 =", v3, reflect.TypeOf(v3))

	simple_log("Leave:\n")
}

// 2.3.5 字符串
func chap02_t08() {
	simple_log("Enter:")

	var v1 string = "我爱中华"
	fmt.Println("v1[0] =", v1[0], len(v1), reflect.TypeOf(v1[0]))

	var v2 string = "abc"
	fmt.Println("v1[0] =", v2[0], reflect.TypeOf(v2[0]))
	simple_log("Leave:\n")
}

// 动态增加元素
func chap02_t09() {
	simple_log("Enter:")

	var v1 []int = []int{1, 2, 3}
	var v2 []int = []int{11, 22, 33}
	v1 = append(v1, v2...)

	fmt.Println(v1)
	fmt.Println(v2)

	simple_log("Leave:\n")

}

// 2.3.8 数组切片
// 基于数组创建数组切片
func chap02_01() {
	simple_log("Enter:")

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Elements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	simple_log("Leave:\n")
}

// 数组切片容量
func chap02_02() {
	simple_log("Enter:\n")
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("len(mySlice):", cap(mySlice))
	simple_log("Leave:\n")
}

func chap02_03() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	var persionDB map[string]PersonInfo
	persionDB = make(map[string]PersonInfo)
	persionDB["12345"] = PersonInfo{"12345", "Tom", "Room 203, ..."}
	persionDB["1"] = PersonInfo{"1", "Jack", "Jack 101, ..."}
	person, ok := persionDB["12345"]

	if ok {
		fmt.Println("Found person", person)

	} else {
		fmt.Println("Did not find person")
	}
}

////////////////////////////////////////////////////////////////////////////////
func chap02_main() {

	chap02_t01()

	chap02_t02()

	chap02_t03()

	chap02_t04()

	chap02_t05()

	chap02_t06()

	chap02_t07()

	chap02_t08()

	chap02_t09()

	chap02_01()

	chap02_02()

	chap02_03()
}
