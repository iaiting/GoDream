// =============================================================================
//
// Author            : iaiting
//
// Contact           : iaiting@aliyun.com
//
// Generate Data     : 2018-02-17
//
// Copyright         : 本文件隶属iaiting，欢迎转载无须知会
//
// Desciption        : 密码学包之对称加密aes加解密模块
//
// =============================================================================
package xcrypto

import (
	"crypto/aes"
)

// =============================================================================
//
// AES 加密相关函数
//
// =============================================================================
func AES128_Encrypt(key [16]byte, in []byte) ([]byte, error) {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	o_bs := make([]byte, len(in))
	c.Encrypt(o_bs, in)

	return o_bs, nil
}

func AES192_Encrypt(key [24]byte, in []byte) ([]byte, error) {
	return nil, nil
}

func AES256_Encrypt(key [32]byte, in []byte) ([]byte, error) {
	return nil, nil
}

// =============================================================================
//
// AES 解密相关函数
//
// =============================================================================
func AES128_Decrypt(key [16]byte, in []byte) ([]byte, error) {
	c, err := aes.NewCipher(key[:])
	if err != nil {
	}

	m_bs := make([]byte, len(in))
	c.Decrypt(m_bs, in)

	return m_bs, nil
}

func AES192_Decrypt(key [24]byte, in []byte) ([]byte, error) {
	return nil, nil
}

func AES256_Decrypt(key [32]byte, in []byte) ([]byte, error) {
	return nil, nil
}
