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
// 以下为自行练习例子
//
// =============================================================================

func chap02_01() {
	fmt.Println("sssssssss")
}

// =============================================================================
//
// 以下为书本原生例子
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
