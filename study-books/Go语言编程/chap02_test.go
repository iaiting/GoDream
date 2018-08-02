package goprogram

import (
	"fmt"
	"testing"
)

// =============================================================================
//
// 以下为自行练习的例子测试程序
//
// =============================================================================
func Test_chap0201_01(t *testing.T) {
	chap0201_01()
}

func Test_chap0201_02(t *testing.T) {
	chap0201_02()
}

func Test_chap0201_03(t *testing.T) {
	chap0201_03()
}
func Test_chap0201_04(t *testing.T) {
	chap0201_04()
}

func Test_chap0202_01(t *testing.T) {
	chap0202_01()
}

func Test_chap0202_02(t *testing.T) {
	chap0202_02()
}

func Test_chap0203_01(t *testing.T) {
	chap0203_01()
}

func Test_chap0203_02(t *testing.T) {
	chap0203_02()
}

func Test_chap0203_03(t *testing.T) {
	t.Log(chap0203_03(3.12345, 3.1236, 0.00001))
}

func Test_chap0203_04(t *testing.T) {
	chap0203_04()
}

func Test_chap0203_05(t *testing.T) {
	chap0203_05()
}

func Test_chap0203_06(t *testing.T) {
	chap0203_06()
}

func Test_chap0203_07(t *testing.T) {
	chap0203_07()
}

func Test_chap0203_08(t *testing.T) {
	chap0203_08()
}
func Test_chap0203_09(t *testing.T) {
	chap0203_09()
}

func Test_chap0203_10(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	chap0203_10(a)
	t.Log("in invoke array values:", a)
}

func Test_chap0203_11(t *testing.T) {
	chap0203_11()
}

func Test_chap0203_12(t *testing.T) {
	chap0203_12()
}

func Test_chap0203_13(t *testing.T) {
	chap0203_13()
}

func Test_chap0203_14(t *testing.T) {
	chap0203_14()
}

func Test_chap0204_01(t *testing.T) {
	fmt.Println(chap0204_01(0))
	fmt.Println(chap0204_01(3))
}

func Test_chap0204_02(t *testing.T) {
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		chap0204_02(v, v)
	}
}

func Test_chap0204_03(t *testing.T) {
	sum1, sum2 := chap0204_03(10)
	t.Log(sum1, sum2)
}

func Test_chap0204_04(t *testing.T) {
	chap0204_04()
}

func Test_chap0205_01(t *testing.T) {
	chap0205_01(1, 2, 2, 1, 3, 4, 4, 3)
	chap0205_01(1, 3, 5)
}
func Test_chap0205_02(t *testing.T) {
	chap0205_02(10, 20, 21, 11, 31, 41, 41, 31)
}

func Test_chap0205_03(t *testing.T) {
	chap0205_03()
}
func Test_chap0205_04(t *testing.T) {
	chap0205_041()
}

// =============================================================================
//
// 以下为书本原生例子的测试程序
//
// =============================================================================
func Test_Chap02_01(t *testing.T) {
	Chap02_01()
}

func Test_Chap02_02(t *testing.T) {
	Chap02_02()
}

func Test_Chap02_03(t *testing.T) {
	Chap02_03()
}

func Test_Chap02_04(t *testing.T) {
	Chap02_04(1, "qweabc", 44)
	Chap02_04(1, "qweabc", 44, 12.3)
}
