package xasn1

import (
	"fmt"
	"encoding/hex"
	"encoding/asn1"
)

func Hex2bytes(s string) (b []byte, e error) {
	bb, _:= hex.DecodeString(s)
	fmt.Println(bb)
	return nil, nil
}

func Bytes2hex(b [] byte) (s string, e error) {
	ss:= hex.EncodeToString(b)
	fmt.Println(ss)
	return ss, nil
}

type SysPid struct {
	Version 	int
	SysID 		[]byte
	SysPK 		[]byte
}

func ParseSysPid(syspidHex string) (e error) {
	syspidBytes, _ := hex.DecodeString(syspidHex)
	syspid := new(SysPid)
	asn1.Unmarshal(syspidBytes, syspid)

	fmt.Println("============")
	fmt.Println(syspid.Version)
	fmt.Println("============")
	fmt.Println(string(syspid.SysID[:]))
	fmt.Println("============")
	for i:=0;i<len(syspid.SysPK);i++ {
		fmt.Printf("%02x ",syspid.SysPK[i])
		if((i-31)%32==0) {
			fmt.Printf("\012")
		}
	}
	return nil
}


type UserPid struct {
    Version      int     `asn1:"default:1"`
    Sysid        asn1.RawValue `asn1:"tag:0,optional,explicit"`
    ID           []byte
    Pid          []byte
}

func ParseUserPid(syspidHex string) (e error) {
	userpidBytes, _ := hex.DecodeString(syspidHex)
	userpid_st := new(UserPid)
	asn1.Unmarshal(userpidBytes, userpid_st)

	fmt.Println("============")
	fmt.Println(userpid_st.Version)
	fmt.Println("============")
	fmt.Println(string(userpid_st.ID[:]))

	// fmt.Println("============")
	// for i:=0;i<len(syspid.SysPK);i++ {
	// 	fmt.Printf("%02x ",syspid.SysPK[i])
	// 	if((i-31)%32==0) {
	// 		fmt.Printf("\012")
	// 	}
	// }
	return nil
}