package cms

import (
	"encoding/hex"
	"crypto/x509"
	"log"
)

////////////////////////////////////////////////////////////////////////////////
func getSerialNumberFromCertAsn1Bytes(certAsn1Bytes []byte) string {
	pCertificate, _ := x509.ParseCertificate(certAsn1Bytes)
	snBytes := pCertificate.SerialNumber.Bytes()
	snHexstring :=hex.EncodeToString(snBytes)
	return snHexstring
}

func GetSerialNumberFromCertAsn1Bytes(certAsn1Bytes []byte) string {
	return getSerialNumberFromCertAsn1Bytes(certAsn1Bytes)
}

func GetSerialNumberFromCertAsn1HexString(certAsn1Hexstring string) string {
	certAsn1Bytes, _ := hex.DecodeString(certAsn1Hexstring)
	return getSerialNumberFromCertAsn1Bytes(certAsn1Bytes)
}
////////////////////////////////////////////////////////////////////////////////
func getCommonNameFromCertAsn1Bytes(certAsn1Bytes []byte) string {
	pCertificate, _ := x509.ParseCertificate(certAsn1Bytes)
	cnString := pCertificate.Issuer.CommonName
	return cnString
}

func GetCommonNameFromCertAsn1Bytes(certAsn1Bytes []byte) string {
	return getCommonNameFromCertAsn1Bytes(certAsn1Bytes)
}

func GetCommonNameFromCertAsn1HexString(certAsn1Hexstring string) string {
	certAsn1Bytes, _ := hex.DecodeString(certAsn1Hexstring)
	return getCommonNameFromCertAsn1Bytes(certAsn1Bytes)
}
////////////////////////////////////////////////////////////////////////////////
// func getCommonNameFromCertAsn1Bytes(certAsn1Bytes []byte) string {
// 	pCertificate, _ := x509.ParseCertificate(certAsn1Bytes)
// 	cnString := pCertificate.Issuer.CommonName
// 	return cnString
// }

// func GetCommonNameFromCertAsn1Bytes(certAsn1Bytes []byte) string {
// 	return getCommonNameFromCertAsn1Bytes(certAsn1Bytes)
// }

func GetIssuerFromCertAsn1HexString(certAsn1Hexstring string) {
	certAsn1Bytes, _ := hex.DecodeString(certAsn1Hexstring)
	pCertificate, _ := x509.ParseCertificate(certAsn1Bytes)
	// pCertificate.Issuer.
	log.Println("****",pCertificate.Issuer.Country)
	log.Println("****",pCertificate.Issuer.Organization)
	log.Println("****",pCertificate.Issuer.CommonName)
	// fmt.pCertificate.Issuer
	// return getCommonNameFromCertAsn1Bytes(certAsn1Bytes)
}

////////////////////////////////////////////////////////////////////////////////
func GetIssuerFromCertAsn1HexString33(certAsn1Hexstring string) {
	certAsn1Bytes, _ := hex.DecodeString(certAsn1Hexstring)
	pCertificate, _ := x509.ParseCertificate(certAsn1Bytes)
	// pCertificate.Issuer.
	log.Println("****",pCertificate.EmailAddresses)
	log.Println("****",pCertificate.RawIssuer)
	log.Println("****",pCertificate.RawSubject)
}
