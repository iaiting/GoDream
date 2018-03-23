package recipient

import (
	"fmt"
	"encoding/hex"
	"crypto/x509"
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getSerialNumber(certDerBytes []byte) (string, error) {
	pX509, err := x509.ParseCertificate(certDerBytes)
	if err != nil {
		      return "", nil
	}

	serialNumberStr := pX509.SerialNumber.String()

	return serialNumberStr, nil	
}

func GetSerialNumberFromDerBytes(certDerBytes []byte) (string, error) {
	return getSerialNumber(certDerBytes)
}

func GetSerialNumberFromDerHex(certDerHex string) (string, error) {
	bs, err := hex.DecodeString(certDerHex)
	if err != nil {
		return "", nil
	}
	return getSerialNumber(bs)
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getIssuerName(certDerBytes []byte) (string, error) {
	pX509, err := x509.ParseCertificate(certDerBytes)
	if err != nil {
		return "", nil
	}
	cn := pX509.Issuer.CommonName
	c := pX509.Issuer.Country[0]
	o := pX509.Issuer.Organization[0]
	issuerNameStr := "CN=" + cn + "," + "O=" + o + "," + "C=" + c
	return issuerNameStr, nil	
}

func GetIssuerNameFromDerBytes(certDerBytes []byte) (string, error) {
	return getIssuerName(certDerBytes)
}

func GetIssuerNameFromDerHex(certDerHex string) (string, error) {
	bs, err := hex.DecodeString(certDerHex)
	if err != nil {
		return "", nil
	}
	return getIssuerName(bs)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getSubjectName(certDerBytes []byte) (string, error) {
	pX509, err := x509.ParseCertificate(certDerBytes)
	if err != nil {
		return "", nil
	}

	email := "email="+pX509.EmailAddresses[0]
	cn := "CN="+pX509.Subject.CommonName
	l := "L="+pX509.Subject.Locality[0]
	p := "P="+pX509.Subject.Province[0]
	c := "C="+pX509.Subject.Country[0]

	subjectNameStr := email + "," + cn + "," + l + "," + p + "," + c

	return subjectNameStr, nil	
}

func GetSubjectNameFromDerHex(certDerHex string) (string, error) {
	bs, err := hex.DecodeString(certDerHex)
	if err != nil {
		return "", nil
	}
	return getSubjectName(bs)
}

// 获取证书的类型
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func ri_gettype(recipientDerBytes []byte) (int, error) {
	pX509, err := x509.ParseCertificate(recipientDerBytes)
	if err != nil {
		return -1, nil
	}

	fmt.Println(pX509.RawSubjectPublicKeyInfo)
	ritype  := pX509.PublicKeyAlgorithm
	fmt.Println(ritype)

	return 1, nil
}

func RI_GettypeFromDerbytes(recipientDerBytes []byte) (int, error) {
	return ri_gettype(recipientDerBytes)
}

func RI_GettypeFromDerhex(recipientDerHex string) (int, error) {
	recipientDerBytes, err := hex.DecodeString(recipientDerHex)
	if err!=nil {
		return -1, err
	}

	return ri_gettype(recipientDerBytes)
}