package cms

import (
	"fmt"
	"crypto/x509/pkix"
	"encoding/asn1"
	// "encoding/base64"
	// base64 "encoding/base64"
	"crypto/rand"
	// "encoding/hex"
	// "fmt"
)
////////////////////////////////////////////////////////////////////////////////
var (
	dataOid = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}
	signedDataOid = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 2}
	envelopedDataOid = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 3}
)

////////////////////////////////////////////////////////////////////////////////
type CMS struct {
}

////////////////////////////////////////////////////////////////////////////////
type contentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content asn1.RawValue
}

////////////////////////////////////////////////////////////////////////////////
type recipientInfo struct {
	Version int
	IssuerAndSerialNumber issuerAndSerial
	KeyEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedKey []byte
}

////////////////////////////////////////////////////////////////////////////////
type encryptedContentInfo struct {
	ContentType asn1.ObjectIdentifier
	ContentEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedContent asn1.RawValue `asn1:"tag:0,optional,explicit"`
}

////////////////////////////////////////////////////////////////////////////////
type issuerAndSerial struct {
}

////////////////////////////////////////////////////////////////////////////////
type signedData struct {
	Version int
	DigestAlgorithms []pkix.AlgorithmIdentifier
}

////////////////////////////////////////////////////////////////////////////////
type envelopedData struct {
	Version int
	RecipientInfos []recipientInfo
	EncryptedContentInfo encryptedContentInfo
}


var (
	// 非对称算法oid
	oidDigestAlgorithmSHA1    = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 26}
	oidEncryptionAlgorithmRSA = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}

	// 对称算法oid
	oidEncryptionAlgorithmDESCBC = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 7}
)

func marshalEncryptedContent(contentBytes []byte) asn1.RawValue {
	asn1ContentBytes, _ := asn1.Marshal(contentBytes)
	return asn1.RawValue {
		Tag: 0,
		Class: 2,
		Bytes: asn1ContentBytes,
		IsCompound: true,
	}
}
func myEncrypt(inBytes []byte) ([]byte, *encryptedContentInfo, error) {
	var err error
	key := make([]byte, 32)
	iv := make([]byte, 16)
	rand.Read(key)
	rand.Read(iv)

	eci:= encryptedContentInfo {
		ContentType:dataOid,
		ContentEncryptionAlgorithm: pkix.AlgorithmIdentifier {
			Algorithm:oidEncryptionAlgorithmDESCBC,
		},
		EncryptedContent:marshalEncryptedContent([]byte("123")),
	}
	return key, &eci, err
}

func MyEncrypt(inBytes []byte, okeyIn []byte) (okey []byte, outBytes []byte, err error) {
	okey = make([]byte, 32)
	if okeyIn == nil {		
		rand.Read(okey)
	} else {
		okey = okeyIn
	}

	// 对 inBytes 加密
	outBytes = inBytes

	// error
	err = nil
	return 
}
////////////////////////////////////////////////////////////////////////////////
func CmsEncrypt(inBytes []byte, certHexstrings []string)([]byte, error) {
	// okey, eci, _ := myEncrypt(inBytes)

	// var tt []byte 
	var t []byte 
	for i:=0;i<100;i++ {
		// t = make([]byte, 1024 *1024 *100)
		t = make([]byte, 1024 *100)
		// tt += t
	}
	
	// fmt.Println("******",t, len(t))
	fmt.Println("******", len(t))
	
	_, _,_ = MyEncrypt(inBytes, nil)
	return nil, nil

	// ris := make([]recipientInfo, len(certHexstrings))
	// for n, certHexstring := range certHexstrings {
	// 	fmt.Println("Once certHexstring", certHexstring)
	// 	ek := []byte("0123456789")
	// 	ri := recipientInfo {
	// 		Version: 0,
	// 		IssuerAndSerialNumber :issuerAndSerial{},
	// 		KeyEncryptionAlgorithm: pkix.AlgorithmIdentifier {
	// 			Algorithm: oidEncryptionAlgorithmRSA,
	// 		},
	// 		EncryptedKey: ek,
	// 	}
	// 	ris[n] = ri
	// }

	// envdata := envelopedData {
	// 	Version: 0,
	// 	RecipientInfos: ris,
	// 	EncryptedContentInfo: *eci,
	// }
	// envdataAsn1, err := asn1.Marshal(envdata)

	// fmt.Println("rv1 = ", envdataAsn1,err)
	// wrapper := contentInfo {
	// 	ContentType: envelopedDataOid,
	// 	Content: asn1.RawValue{Class: 2, Tag: 0, IsCompound: true, Bytes: envdataAsn1},
	// }
	// rv, err:= asn1.Marshal(wrapper)
	// fmt.Println("rv = ", rv,err)
	// return rv, err
}

////////////////////////////////////////////////////////////////////////////////
func CmsDecrypt()([]byte, error) {
	return nil,nil
}