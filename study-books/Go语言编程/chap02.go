// =============================================================================
//
// Author            : iaiting
//
// Contact           : iaiting@aliyun.com
//
// Generate Data     : 2018-02-17
//
// Copyright         : 本文件隶属iaiting，欢迎转载无须知会
//
// Desciption        : Go语言编程(许式伟 吕桂华)第2章练习题
//                     大写字母开头的为书本例题编号，小写字母开头的为自行连续编号
//
// =============================================================================
package goprogram

import (
	"fmt"
)

// =============================================================================
//
// 以下为自行练习的例子
//
// =============================================================================

// 变量定义
// =============================================================================
func chap0201_01() {
	var int_a int
	int_a = 123

	var string_a string
	string_a = "abc123"

	var aint_a [10]int
	aint_a = [10]int{1, 2, 3}

	var sint_a []int
	sint_a = []int{3, 2, 1}

	var struct_a struct {
		int_a    int
		string_a string
		aint_a   [10]int
		sint_a   []int
	}

	var pint_a *int
	pint_a = &int_a

	var msringint_a map[string]int

	// var func_a func(a int) int

	var (
		int_b    int
		string_b string
	)
	int_b = 123
	string_b = "abc"

	// 打印变量
	fmt.Println(int_a)
	fmt.Println(string_a)
	fmt.Println(aint_a)
	fmt.Println(sint_a)
	fmt.Println(struct_a)
	fmt.Println(pint_a)
	fmt.Println(msringint_a)
	// fmt.Println(func_a(1))
	fmt.Println(int_b)
	fmt.Println(string_b)
}

// 变量初始化
// =============================================================================
func chap0201_02() {
	var int_a1 int = 1234567890
	var int_a2 = 2345678901
	int_a3 := 3456789012

	fmt.Println(int_a1)
	fmt.Println(int_a2)
	fmt.Println(int_a3)
}

// 变量赋值
// =========================================================================
func chap0201_03() {
	var int_a int
	int_a = 1234567890

	var int_b int
	int_b = 2345678901

	fmt.Println(int_a, int_b)

	int_a, int_b = int_b, int_a
	fmt.Println(int_a, int_b)
}

// 匿名变量
// =============================================================================
func chap0201_04() {
	_, _, nickName := "Qiao", "ting", "iaiting"
	fmt.Println(nickName)
}

// =============================================================================
//
// 以下为书本原生的例子
//
// =============================================================================

func Chap02_01() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	for i, v := range myArray {
		fmt.Println(i, ":", v)
	}

	fmt.Println("--------------------")
	for i, v := range mySlice {
		fmt.Println(i, ":", v)
	}
}

// =============================================================================
func Chap02_02() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}

// =============================================================================
func Chap02_03() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	// var personDB map[string]PersonInfo
	personDB := make(map[string]PersonInfo)
	personDB["123"] = PersonInfo{"123", "tom", "mas"}
	personDB["1"] = PersonInfo{"1", "Jack", "nj"}

	person, ok := personDB["123"]
	if ok {
		fmt.Println("found ok", person)
	} else {
		fmt.Println("found not ok")
	}
}

// =============================================================================
