package xcrypto

import (
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm2"
)

// =============================================================================
//
// SM2 公私钥相关函数
//
// =============================================================================
func SM2_KeyGenerate() (*sm2.PrivateKey, *sm2.PublicKey) {
	prikey, err := sm2.GenerateKey()
	if err != nil {
	}
	pubkey := &prikey.PublicKey

	return prikey, pubkey
}

// =============================================================================
//
// SM2 加密相关函数
//
// =============================================================================
func SM2_Encrypt(pubkey *sm2.PublicKey, in []byte) ([]byte, error) {
	return pubkey.Encrypt(in)
}

func SM2_EncryptBytes2Hexstr(pubkey *sm2.PublicKey, inBytes []byte) (string, error) {
	bs, err := SM2_Encrypt(pubkey, inBytes)
	if err != nil {

	}

	return hex.EncodeToString(bs), nil
}

// =============================================================================
//
// SM2 解密相关函数
//
// =============================================================================
func SM2_Decrypt(prikey *sm2.PrivateKey, in []byte) ([]byte, error) {
	return prikey.Decrypt(in)
}

func SM2_DecryptHexstr2Bytes(prikey *sm2.PrivateKey, inHexstr string) ([]byte, error) {
	bs, err := hex.DecodeString(inHexstr)
	if err != nil {
	}

	return SM2_Decrypt(prikey, bs)
}

// =============================================================================
//
// SM2 签名相关函数
//
// =============================================================================
func SM2_Sign() {
	return
}

// =============================================================================
//
// SM2 验签相关函数
//
// =============================================================================
func SM2_Verify() {
	return
}
