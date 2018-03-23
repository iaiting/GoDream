package xcrypto

import (
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
//
// SM2 测试
//
////////////////////////////////////////////////////////////////////////////////

func Test_SM2(t *testing.T) {
	t.Log("**************\n")
	prikey, pubkey := SM2_KeyGenerate()

	in := []byte("abc123")
	o, _ := SM2_EncryptBytes2Hexstr(pubkey, in)
	t.Log("**************", o)

	d_o, _ := SM2_DecryptHexstr2Bytes(prikey, o)

	t.Log("**************", d_o)
}
