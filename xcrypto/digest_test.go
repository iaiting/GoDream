package xcrypto

import (
	"fmt"
	"testing"
)

func Test_MD5_Encryptstring(t *testing.T) {
	hhu := MD5_Encryptstring("abc123")
	t.Log(hhu)
	fmt.Println(hhu)
}

func Test_MD5_Encryptbytes(t *testing.T) {
	hhu := MD5_Encryptbytes([]byte("abc123"))
	t.Log(hhu)
	fmt.Println(hhu)
}
