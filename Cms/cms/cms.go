package cms

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
)

type contentType_st struct {
	ContentType asn1.ObjectIdentifier
}

type contentInfo_st struct {
	ContentType contentType_st
	content     asn1.RawValue
}

type signedData_st struct {
	version int
}

/////////////////////////////////////////////////////////////////////////////////////////
type originatorInfo_st struct {
}

/////////////////////////////////////////////////////////////////////////////////////////
type keyTransRecipientInfo_st struct {
}

type keyAgreeRecipientInfo_st struct {
}

type kEKRecipientInfo_st struct {
}

type passwordRecipientinfo_st struct {
}

type otherRecipientInfo_st struct {
}

type recipientInfo_st struct {
	ktri  keyTransRecipientInfo_st `asn1:"optional"`
	kari  keyAgreeRecipientInfo_st `asn1:"tag:1,optional"`
	kekri kEKRecipientInfo_st      `asn1:"tag:2,optional"`
	pwri  passwordRecipientinfo_st `asn1:"tag:3,optional"`
	ori   otherRecipientInfo_st    `asn1:"tag:4,optional"`
}

/////////////////////////////////////////////////////////////////////////////////////////
type encryptedContentInfo_st struct {
	ContentType                asn1.ObjectIdentifier
	ContentEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedContent.RawValue  `asn1:"tag:0,explicit,optional"`
}

func NewEncryptedContentInfo(ct []int, alg []int) (ec *encryptedContentInfo_st) {
	return &encryptedContentInfo_st{
		ContentType: ct,
		ContentEncryptionAlgorithm: pkix.AlgorithmIdentifier{
			Algorithm: alg,
		},
	}
}

/////////////////////////////////////////////////////////////////////////////////////////
type unprotectedAttrs_st struct {
}

/////////////////////////////////////////////////////////////////////////////////////////
type envelopedData_st struct {
	version              int
	originatorInfo       originatorInfo_st  `asn1:"tag:0,optional,explicit"`
	recipientInfos       []recipientInfo_st `asn1:"set"`
	encryptedContentInfo encryptedContentInfo_st
	unprotectedAttrs     unprotectedAttrs_st `asn1:"tag:1,optional,explicit"`
}

/////////////////////////////////////////////////////////////////////////////////////////
func CmsEncrypt(in []byte) (out []byte, err error) {
	return nil, nil
}
func CmsDecrytp(in []byte) (out []byte, err error) {
	return nil, nil
}

func Encrypt(alg int, key []byte, in []byte) (out []byte, err error) {
	fmt.Println("Input Encrypt Args: ")
	fmt.Println("alg = ", alg)
	fmt.Println("key = ", key)
	fmt.Println("in = ", in)
	return in, nil
}
func Decrypt(alg int, key []byte, in []byte) (out []byte, err error) {
	fmt.Println("Input Decrypt Args: ")
	fmt.Println("alg = ", alg)
	fmt.Println("key = ", key)
	fmt.Println("in = ", in)

	return in, nil
}

func SM4Encrypt(in []byte, algID []int) {
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 'a', 'b', 'c', 'd', 'e', 'f'}
	asn1Bytes, _ := asn1.Marshal(in)
	eci := encryptedContentInfo_st{
		ContentType: oidData,
		ContentEncryptionAlgorithm: pkix.AlgorithmIdentifier{
			Algorithm: algID,
			Parameters: asn1.RawValue{
				Tag:   asn1.TagOctetString,
				Bytes: iv,
			},
		},
		EncryptedContent: asn1.RawValue{Tag: 0, Class: 2, Bytes: asn1Bytes, IsCompound: true},
	}

	asn1_bytes, err := asn1.Marshal(eci)
	fmt.Println(asn1_bytes, err)
	hexStr := hex.EncodeToString(asn1_bytes)
	fmt.Println(hexStr)
}
