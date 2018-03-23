package xcrypto

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5_Encryptstring(str string) string {
	h := md5.New()
	_, e := io.WriteString(h, str)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%X", h.Sum(nil))
}

func MD5_Encryptbytes(bytes []byte) string {
	h := md5.Sum(bytes)
	return fmt.Sprintf("%X", h)
}

func SM3_Encrypt() {
}
