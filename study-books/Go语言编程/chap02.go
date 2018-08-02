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
// =============================================================================
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
	_, _, nickName := "Qiao", "Ting", "iaiting"
	fmt.Println(nickName)
}

// 常量
// =============================================================================
func chap0202_01() {
	const PI float64 = 3.1415926
	const ZERO = 0.0
	const (
		SIZE int64 = 1024
		EOF        = -1
	)
	const U, V1 float32 = 0, 3
	const A, B, C = 3, 4, "abc123"
	const MASK = 1 << 3

	fmt.Println(PI)
	fmt.Println(ZERO)
	fmt.Println(SIZE)
	fmt.Println(U, V1)
	fmt.Println(A, B, C)
	fmt.Println(MASK)
}

// 常量
// =============================================================================
func chap0202_02() {
	const (
		CONST0 = iota
		CONST1 = iota
		CONST2 = iota
	)
	fmt.Println(CONST0)
	fmt.Println(CONST1)
	fmt.Println(CONST2)

	const (
		CONSTA = 1 << iota
		CONSTB = 1 << iota
		CONSTC = 1 << iota
	)
	fmt.Println(CONSTA)
	fmt.Println(CONSTB)
	fmt.Println(CONSTC)

	const (
		u  float64 = iota * 42
		v1 float64 = iota * 42.0
		w          = iota * 42.0
	)
	fmt.Println(u)
	fmt.Println(v1)
	fmt.Println(w)

	const x = iota
	const y = iota
	fmt.Println(x)
	fmt.Println(y)

	const (
		C0 = iota
		C1
		C2
	)
	fmt.Println(C0)
	fmt.Println(C1)
}

// 枚举
// =============================================================================
func chap0202_03() {
	const (
		Sunday, Sun = iota, iota
		Monday, Mon
		Tuesday, Tue
		Wednesday, Wed
		Thursday, Thur
		Friday, Fri
		Saturday, Sat
		daysOfWeek, _
	)
	fmt.Println(Sunday, Sun)
	fmt.Println(Monday, Mon)
	fmt.Println(Tuesday, Tue)
	fmt.Println(Wednesday, Wed)

	fmt.Println(Thursday, Thur)
	fmt.Println(Friday, Fri)
	fmt.Println(Saturday, Sat)
	fmt.Println(daysOfWeek)
}

// 布尔类型
// =============================================================================
func chap0203_01() {
	var bool_v11 bool
	bool_v11 = true
	bool_v12 := (1 == 2)

	fmt.Println(bool_v11)
	fmt.Println(bool_v12)
}

// 整数类型
// =============================================================================
func chap0203_02() {
	fmt.Println(5 % 3)
}

// 浮点数
// =============================================================================
func chap0203_03(float_v11, float_v12, p float32) bool {
	return -p <= float_v11-float_v12 && float_v11-float_v12 <= p
}

// 复数
// =============================================================================
func chap0203_04() {
	var complex64_v11 complex64
	complex64_v11 = 3.2 + 12i
	complex64_v12 := 4.3 + 13i
	complex64_v13 := complex(5.4, 14)

	fmt.Println(complex64_v11)
	fmt.Println(complex64_v12)
	fmt.Println(complex64_v13)
	fmt.Println(real(complex64_v13))
	fmt.Println(imag(complex64_v13))
}

// 字符串
// =============================================================================
func chap0203_05() {
	var string_v11 string
	string_v11 = "abc123"
	byte_v11 := string_v11[0]
	fmt.Printf("The length of \"%s\" is %d\n", string_v11, len(string_v11))
	fmt.Printf("The first character of \"%s\" is %c\n", string_v11, byte_v11)
}

// 字符串的字节遍历
// =============================================================================
func chap0203_06() {
	string_v11 := "Hello, 世界"
	n := len(string_v11)
	for i := 0; i < n; i++ {
		byte_ch := string_v11[i]
		fmt.Println(i, ":", byte_ch)
	}
}

// 字符串的unicode字符遍历
// =============================================================================
func chap0203_07() {
	string_v11 := "Hello, 世界"
	for i, ch := range string_v11 {
		fmt.Println(i, ch)
	}
}

// 数组
// =============================================================================
func chap0203_08() {
	var byte_32_v11 [32]byte
	var pfloat32_16_v11 [16]*float32
	var int_3l5c [3][5]int
	var float64_2_3_4 [2][3][4]float64

	fmt.Println(len(byte_32_v11))
	fmt.Println(len(pfloat32_16_v11))
	fmt.Println(len(int_3l5c))
	fmt.Println(len(float64_2_3_4))
}

// 数组遍历
// =============================================================================
func chap0203_09() {
	var byte_32c_v11 [32]byte
	for i := 0; i < len(byte_32c_v11); i++ {
		fmt.Println("Element", i, "of array is", byte_32c_v11[i])
	}

	for i, v1 := range byte_32c_v11 {
		fmt.Println("Array element[", i, "] = ", v1)
	}
	fmt.Println(len(byte_32c_v11))
}

// 数组值参数值传递
// =============================================================================
func chap0203_10(int_c5 [5]int) {
	int_c5[0] = 10
	fmt.Println("in modify array v1alues:", int_c5)
}

// 数组切片创建的三种方式
// =============================================================================
func chap0203_11() {
	// make 方式创建
	intslice_v11 := make([]int, 5)

	// 直接创建
	intslice_v12 := []int{0, 0, 0, 0, 0}

	// 通过数组创建
	intc5_v13 := [5]int{}
	intslice_v13 := intc5_v13[:]

	fmt.Println(intslice_v11, intslice_v12, intslice_v13)

	// make 方式创建
	intslice_v14 := make([]int, 5, 10)
	for i := 0; i < len(intslice_v14); i++ {
		fmt.Println(i, intslice_v14[i])
	}

	for i, v1 := range intslice_v14 {
		fmt.Println(i, v1)
	}
}

// 切片创建切片
// =============================================================================
func chap0203_12() {
	intslice_v11 := make([]int, 5, 5)
	intslice_v12 := intslice_v11[:5]
	fmt.Println(intslice_v11, intslice_v12)
}

// 切片复制
// =============================================================================
func chap0203_13() {
	intslice_v11 := []int{1, 2, 3, 4, 5}
	intslice_v12 := []int{5, 4}
	fmt.Println(intslice_v11, intslice_v12)
	copy(intslice_v11, intslice_v12)
	fmt.Println(intslice_v11, intslice_v12)

	intslice_v13 := []int{9, 8}
	copy(intslice_v13, intslice_v11)
	fmt.Println(intslice_v11, intslice_v13)
}

// 删除map中的元素
// =============================================================================
func chap0203_14() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}

	personDB := make(map[string]PersonInfo, 20)
	personDB["123456"] = PersonInfo{"123456", "abc123", "beijing"}
	personDB["234567"] = PersonInfo{"234567", "bcd234", "shanghai"}
	fmt.Println(len(personDB), personDB)

	delete(personDB, "123456")
	fmt.Println(len(personDB), personDB)

	delete(personDB, "1234567")
	fmt.Println(len(personDB), personDB)

	delete(personDB, "1234567")
	fmt.Println(len(personDB), personDB)
}

// if条件语句
// =============================================================================
func chap0204_01(x int) int {
	if x == 0 {
		return x + 1
	} else {
		return x
	}
}

// switch选择语句
// =============================================================================
func chap0204_02(v1 int, v2 int) {
	switch v1 {
	case 0:
		fmt.Println("0")

	case 1:
		fmt.Println("1")

	case 2:
		fallthrough

	case 3:
		fmt.Println("3")

	case 4, 5, 6:
		fmt.Println("{4,5,6}")

	default:
		fmt.Println("default")
	}

	// 没有条件的switch等同于多个if
	switch {
	case 0 <= v2 && v2 <= 3:
		fmt.Println("0 - 3")

	case 4 <= v2 && v2 <= 6:
		fmt.Println("4 - 6")

	case 7 <= v2:
		fmt.Println("7 - *")
	}
}

// for循环语句
// =============================================================================
func chap0204_03(x int) (int, int) {
	int_sum1 := 0
	for i := 0; i < x; i++ {
		int_sum1 += i
	}

	int_sum2 := 0
	for {
		int_sum2++
		if int_sum2 > 100 {
			break
		}
	}
	return int_sum1, int_sum2
}

// goto跳转语句
// =============================================================================
func chap0204_04() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i > 20 {
			goto END
		}
	}

END:
	fmt.Println("Function END")
}

// 不定参数函数类型
// =============================================================================
func chap0205_01(args ...int) {
	for i, v := range args {
		fmt.Println(i, v)
	}
}

// 不定参数函数调用
// 可以将args理解成切片，args...理解成多个参数
// =============================================================================
func chap0205_02(args ...int) {
	chap0205_01(args...)
	chap0205_01([]int{2, 3, 4000, 22}[2:]...)
}

// 匿名函数
// =============================================================================
func chap0205_03() {
	sum := func(v1, v2 int) int {
		return v1 + v2
	}(2, 3)
	println("v1 + v2 =================", sum)
}

// 闭包
// =============================================================================
func chap0205_04() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			println("i = %d, j= %d\012", i, j)
		}
	}() // 直接调用返回一个函数指针，变量i可以保证线程安全性，只有匿名函数内部可以访问

	a()
	j = j * 2
	a()
}

// 错误处理，自定义error接口
// =============================================================================
type MyError struct {
	Op   string
	Path string
	err  error
}

func (me *MyError) Error() string {
	return me.Op + me.Path + me.err.Error()
}

// defer
// =============================================================================
func CopyFile(dst, src string) (r int64, e error) {
	defer func() {
		println("This is defer")
	}()

	return
}

// defer
// =============================================================================
func chap0205_05(dst, src string) (r int64, e error) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	return
}

// =============================================================================
//
// 以下为书本原生例子
//
// =============================================================================

// 切片
// =============================================================================
func Chap02_01() {
	var intc10_v11 [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var intslice_v11 []int = intc10_v11[:5]

	fmt.Println("Elements of intc10_v11:")
	for _, v1 := range intc10_v11 {
		fmt.Print(v1, " ")
	}

	fmt.Println("\nElements of intslice_v11:")
	for _, v1 := range intslice_v11 {
		fmt.Print(v1, " ")
	}
}

// 切片容量
// =============================================================================
func Chap02_02() {
	intslice := make([]int, 5, 10)
	fmt.Println("len(intslice):", len(intslice))
	fmt.Println("cap(intslice):", cap(intslice))

	fmt.Println(intslice)
	// intslicet := append(intslice, 1, 2, 3, 4, 5, 6, 7, 8)
	intslicet := append(intslice, 1, 2)
	fmt.Println(intslice)
	fmt.Println(intslicet)
}

// map字典数据结构
// =============================================================================
func Chap02_03() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}

	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["123456"] = PersonInfo{"123456", "abc123", "beijing"}
	personDB["234567"] = PersonInfo{"234567", "bcd234", "shanghai"}

	person, ok := personDB["123456"]
	if ok {
		fmt.Println("Found person", person.Name, "with id", person.ID)
	} else {
		fmt.Println("Did not find person")
	}
}

// ...interface{}任意类型个数不确定参数
// =============================================================================
func Chap02_04(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// =============================================================================
