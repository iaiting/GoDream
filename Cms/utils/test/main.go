package main

import (
	"fmt"
	"GoDream/Cms/utils"
)

func TEST_BYTE_Convert2hexlower_0() {
	d := []byte{0x1E,0x3e,0xE}
	hl := utils.BYTE_Convert2hexlower_0(d)
	fmt.Println(hl)
}

func TEST_BYTE_Convert2hexlower_1() {
	d := []byte{0x1E,0x3e,0xE}
	hl := utils.BYTE_Convert2hexlower_1(d)
	fmt.Println(hl)
}

func TEST_BYTE_Convert2hexupper_0() {
	d := []byte{0x1E,0x3e,0xE}
	hu := utils.BYTE_Convert2hexupper_0(d)
	fmt.Println(hu)
}

func TEST_BYTE_Convert2hexupper_1() {
	d := []byte{0x1E,0x3e,0xE}
	hu := utils.BYTE_Convert2hexupper_1(d)
	fmt.Println(hu)
}

func TEST_HEX_Convert2byte() {
	d := "333038314333303230313031383036323330363030323031303138303034353334343442353330343037374536423637373333333330333130343443303130343946384637344430394645333541433843364242333137443246434238363333413237443439393746353243363335443841353836343838443533323943303439334137354438383543423042313431443030303335303835444336444236363739304530314546434643413731343142434336443635304144323531303533313730373234323730373234463446373830353230343043363133303330333033303330333033303330333033303331303434433030303438343730353738414132334446344531383738334539373341453339393339394135434143393130374333454532333745324643303932323846354246344137373737383539464233333642384635414146363534313045313538443237423735384631333241323442443444453341353936373934324245434446353932313137313230353137303532344639313143434539"
	b, _:= utils.HEX_Convert2byte(d)
	fmt.Println(string(b))
}


func main() {
	// TEST_BYTE_Convert2hexlower_0()
	// TEST_BYTE_Convert2hexlower_1()

	// TEST_BYTE_Convert2hexupper_0()
	// TEST_BYTE_Convert2hexupper_1()

	TEST_HEX_Convert2byte()
}