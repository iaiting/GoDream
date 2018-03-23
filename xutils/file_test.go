package xutils

import (
	"testing"
)

func Test_File(t *testing.T) {
	inFile := "c:\\Projects\\lescpsn\\GoProjects\\src\\GoDream\\testdata\\CentOS-7-x86_64-DVD-1708.iso"
	bs, err := File_Read(inFile)
	if err != nil {
	}
	outFile := "c:\\Projects\\lescpsn\\GoProjects\\src\\GoDream\\testdata\\CentOS-7-x86_64-DVD-1708.iso.o"
	File_Write(outFile, bs)
}
