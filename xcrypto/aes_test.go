package xcrypto

import (
	"testing"
)

func Test_AES(t *testing.T) {
	// key := [16]byte("1234567890abcdef")
	key := [16]byte{0x01, 0x12}
	in := []byte("1234567890abcdef")

	eo_bs, _ := AES128_Encrypt(key, in)

	mo_bs, _ := AES128_Decrypt(key, eo_bs)

	// mo_bs, _ := AES_Decrypt(key, eo_bs)

	t.Log("\n", eo_bs, "\n")
	t.Log("\n", mo_bs, "\n")
}
