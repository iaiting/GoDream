// =============================================================================
//
// Author            : iaiting
//
// Contact           : iaiting@aliyun.com
//
// Generate Data     : 2018-03-30
//
// Copyright         : 本文件隶属iaiting，欢迎转载无须知会
//
// Desciption        : Go语言编程(许式伟 吕桂华) 第3章 练习题
//                     大写字母开头的为书本例题编号，小写字母开头的为自行连续编号
//
// =============================================================================
package goprogram

import (
	"fmt"
)

// =============================================================================

type Integer int

func (i Integer) Less(o Integer) bool {
	return i < o
}

func (i *Integer) Add(o Integer) {
	*i += o
}

// =============================================================================
func chap0301_01(i Integer) {
	if i.Less(4) {
		println(i, "is less 4")
	} else {
		println(i, "is not less 4")
	}
}

// =============================================================================
func chap0301_02(i Integer) {
	var v Integer = 32
	i.Add(v)
	println(i)
}

// 结构体
// =============================================================================
type Rect struct {
	x, y           float64
	width, heighth float64
}

func (r *Rect) Area() float64 {
	return r.width * r.heighth
}

func chap0301_03() {
	r1 := new(Rect)
	r1.width = 5
	r1.heighth = 5.8
	println("Area1 = ", r1.Area())

	r2 := &Rect{}
	println("Area2 = ", r2.Area())

	r3 := &Rect{0, 0, 4, 5}
	println("Area3 = ", r3.Area())

	r4 := &Rect{width: 8, heighth: 9}
	println("Area4 = ", r4.Area())
}

// 接口
// =============================================================================
type IFile interface {
	Read()
	Write()
	Seek()
	Close()
}

type IReader interface {
	Read()
}

type IWriter interface {
	Write()
}

type ICloser interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func (f *File) Write() {
}

func (f *File) Seek() {
}

func (f *File) Close() {
}

func chap0301_04() {
	var file1 IFile = new(File)
	var file2 IReader = new(File)
	var file3 IWriter = new(File)
	var file4 ICloser = new(File)
}

// =============================================================================
//
// 以下为书本原生例子
//
// =============================================================================
func Chap03_01() {
	fmt.Println("**************:Chap03_01")
}

// =============================================================================
