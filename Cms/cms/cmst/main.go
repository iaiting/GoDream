package main
import (
	"GoDream/cms/cms"
	"fmt"
)
var (
	key []byte = []byte{1,2,3,4}
)

func EncryptTest(in []byte) {
	rv, _ := cms.Encrypt(101, key, in)
	fmt.Println(rv,"\012")
}

func DecryptTest(in []byte) {
	rv, _ := cms.Decrypt(101, key, in)

	fmt.Println(rv,"\012")
}

func SM4EncryptTest() {
	cms.SM4Encrypt([]byte{'1', '2', }, cms.OidEncryptionAlgorithmSM4)
	fmt.Println("\012")
}


func main() {
	EncryptTest([]byte("12"))
	DecryptTest([]byte("12"))
	SM4EncryptTest()
}