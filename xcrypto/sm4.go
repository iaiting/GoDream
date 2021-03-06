package xcrypto

import (
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm4"
)

// =============================================================================
//
// SM4 加密相关函数
//
// =============================================================================
func SM4_Encrypt(key [16]byte, in []byte) ([]byte, error) {
	e_bs := make([]byte, len(in))
	sm4.EncryptBlock(key[:], e_bs, in)
	return e_bs, nil
}

// =============================================================================
func SM4_EncryptBytes2Bytes(key [16]byte, in []byte) ([]byte, error) {
	return SM4_Encrypt(key, in)
}

// =============================================================================
func SM4_EncryptBytes2Hexstr(key [16]byte, in []byte) (string, error) {
	bs, _ := SM4_Encrypt(key, in)

	return hex.EncodeToString(bs), nil
}

// =============================================================================
//
// SM4 解密相关函数
//
// =============================================================================
func SM4_Decrypt(key [16]byte, in []byte) ([]byte, error) {
	m_bs := make([]byte, len(in))
	sm4.DecryptBlock(key[:], m_bs, in)
	return m_bs, nil
}

// =============================================================================
func SM4_DecryptBytes2Bytes(key [16]byte, in []byte) ([]byte, error) {
	return SM4_Decrypt(key, in)
}

// =============================================================================
func SM4_DecryptHexstr2Bytes(key [16]byte, in string) ([]byte, error) {
	bs, _ := hex.DecodeString(in)

	return SM4_Decrypt(key, bs)
}
