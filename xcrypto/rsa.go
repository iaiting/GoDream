package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
//
// RSA 相关函数
//
////////////////////////////////////////////////////////////////////////////////
func RSA_KeyGenerate(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {

	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	pubkey := privkey.Public().(*rsa.PublicKey)

	fmt.Println(privkey, "\n", pubkey)
	return privkey, pubkey

}

////////////////////////////////////////////////////////////////////////////////
func RSA_Encrypt(pubkey *rsa.PublicKey, in []byte) ([]byte, error) {

	return rsa.EncryptPKCS1v15(rand.Reader, pubkey, in)

}

////////////////////////////////////////////////////////////////////////////////
func RSA_EncryptBytes2Bytes(pubkey *rsa.PublicKey, inBytes []byte) ([]byte, error) {

	return RSA_Encrypt(pubkey, inBytes)

}

////////////////////////////////////////////////////////////////////////////////
func RSA_EncryptBytes2Hex(pubkey *rsa.PublicKey, inBytes []byte) (string, error) {

	bs, err := RSA_Encrypt(pubkey, inBytes)
	if err != nil {

	}
	return hex.EncodeToString(bs), nil

}

////////////////////////////////////////////////////////////////////////////////
func RSA_EncryptHex2Bytes(pubkey *rsa.PublicKey, inHex string) ([]byte, error) {

	bs, err := hex.DecodeString(inHex)
	if err != nil {

	}

	return RSA_Encrypt(pubkey, bs)

}

////////////////////////////////////////////////////////////////////////////////
func RSA_EncryptHex2Hex(pubkey *rsa.PublicKey, inHex string) (string, error) {

	bs, err := RSA_EncryptHex2Bytes(pubkey, inHex)
	if err != nil {

	}

	return hex.EncodeToString(bs), nil

}

////////////////////////////////////////////////////////////////////////////////
func RSA_Decrypt(prikey *rsa.PrivateKey, in []byte) ([]byte, error) {

	return rsa.DecryptPKCS1v15(rand.Reader, prikey, in)

}

////////////////////////////////////////////////////////////////////////////////
func RSA_DecryptBytes2Bytes(prikey *rsa.PrivateKey, inBytes []byte) ([]byte, error) {

	return RSA_Decrypt(prikey, inBytes)

}

////////////////////////////////////////////////////////////////////////////////
func RSA_DecryptBytes2Hex(prikey *rsa.PrivateKey, inBytes []byte) (string, error) {
	bs, err := RSA_DecryptBytes2Bytes(prikey, inBytes)
	if err != nil {
	}

	return hex.EncodeToString(bs), err
}

////////////////////////////////////////////////////////////////////////////////
func RSA_DecryptHex2Bytes(prikey *rsa.PrivateKey, inHexstr string) ([]byte, error) {
	bs, err := hex.DecodeString(inHexstr)
	if err != nil {
	}

	return RSA_DecryptBytes2Bytes(prikey, bs)
}

////////////////////////////////////////////////////////////////////////////////
func RSA_DecryptHex2Hex(prikey *rsa.PrivateKey, inHexstr string) (string, error) {

	bs, err := RSA_DecryptHex2Bytes(prikey, inHexstr)
	if err != nil {
	}

	return hex.EncodeToString(bs), nil
}

////////////////////////////////////////////////////////////////////////////////
func RSA_Sign() {

}

////////////////////////////////////////////////////////////////////////////////
func RSA_Verify() {

}
