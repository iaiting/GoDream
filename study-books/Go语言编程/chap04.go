// =============================================================================
//
// Author            : iaiting
//
// Contact           : iaiting@aliyun.com
//
// Generate Data     : 2018-04-03
//
// Copyright         : 本文件隶属iaiting，欢迎转载无须知会
//
// Desciption        : Go语言编程(许式伟 吕桂华) 第4章 练习题
//                     大写字母开头的为书本例题编号，小写字母开头的为自行连续编号
//
// =============================================================================
package goprogram

import (
	"fmt"
	"runtime"
	"sync"
)

// =============================================================================

// =============================================================================
func Add(x, y int) {
	z := x + y
	fmt.Println("z = ", z)
}

// 这种协程的写法是有错误的，后续的列子再改进
func chap0403_01() {
	for i := 0; i < 10; i++ {
		go Add(i, i+1)
	}
}

// =============================================================================
//
// 以下为书本原生例子
//
// =============================================================================
// 这种协程的写法是有错误的，后续的列子再改进
func Chap04_01() {
	for i := 0; i < 10; i++ {
		go func(x, y int) {
			z := x + y
			fmt.Println(z)
		}(i, i)
	}
}

// go语言的多线程写法
func Chap04_03() {
	var counter int = 0
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(l *sync.Mutex) {
			l.Lock()
			counter++
			fmt.Println("============ counter:", counter)
			l.Unlock()
		}(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			fmt.Println("10个线程执行完成")
			break
		}
	}
}

// channel的用法
func Chap04_04() {
	// 创建一个int型的chan的切片
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		// 创建每一个chan
		chs[i] = make(chan int)
		go func(count int, ch chan int) {
			println("This is count:", count)
			ch <- count
		}(i, chs[i])
	}

	// 迭代切片
	for _, ch := range chs {
		<-ch
	}
}

// =============================================================================
