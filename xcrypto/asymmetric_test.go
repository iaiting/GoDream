package xcrypto

import (
	"testing"
)

func Test_RSA(t *testing.T) {
	prikey, pubkey := RSA_KeyGenerate(1024)

	p_bs := []byte("abc123")
	encData, _ := RSA_Encrypt(pubkey, p_bs)
	decData, _ := RSA_Decrypt(prikey, encData)
	t.Log("\n", encData, "\n", decData)

	e_hex, _ := RSA_EncryptBytes2Hex(pubkey, p_bs)
	d_bs, _ := RSA_DecryptHex2Bytes(prikey, e_hex)
	t.Log("\n", e_hex, "\n", d_bs)
}
