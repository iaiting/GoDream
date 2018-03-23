package cms

import (
	"encoding/asn1"
)
// 对象标识符		   名称			OID
// rsaEncryption		RSA算法标识	1.2.840.113549.1.1.1
// sha1withRSAEncryption	SHA1的RSA签名	1.2.840.113549.1.1.5
// ECC			ECC算法标识	1.2.840.10045.2.1
// SM2			SM2算法标识	1.2.156.10197.1.301
// SM3WithSM2		SM3的SM2签名	1.2.156.10197.1.501
// sha1withSM2		SHA1的SM2签名	1.2.156.10197.1.502
// sha256withSM2		SHA256的SM2签名	1.2.156.10197.1.503
// sm3withRSAEncryption	SM3的RSA签名	1.2.156.10197.1.504
// commonName		主体名		2.5.4.3
// emailAddress		邮箱		1.2.840.113549.1.9.1
// cRLDistributionPoints	CRL分发点	2.5.29.31
// extKeyUsage		扩展密钥用法	2.5.29.37
// subjectAltName		使用者备用名称	2.5.29.17
// CP			证书策略	2.5.29.32
// clientAuth		客户端认证	1.3.6.1.5.5.7.3.2

// SM2证书的公钥算法是使用ECC算法的Oid标识（1.2.840.10045.2.1），---公钥
// 然后公钥参数使用SM2国密算法的Oid标识（1.2.156.10197.1.301） --签名算法


// cms oid
var ( 
	oidData				= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}
	oidSignedData		= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 2}
	oidEnvelopedData	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 3}
)

// 非对称算法oid
var (
	// SM2-1椭圆曲线数字签名算法
	oidEncryptionAlgorithmSM2_1 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 301, 1}

	// SM2-2椭圆曲线密钥交换协议
	oidEncryptionAlgorithmSM2_2 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 301, 2}

	// SM2-3椭圆曲线加密算法
	oidEncryptionAlgorithmSM2_3 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 301, 3}
)

// 对称算法oid
var (
	oidEncryptionAlgorithmDESCBC = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 7}
	oidEncryptionAlgorithmDESEDE3CBC = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 7}
	oidEncryptionAlgorithmAES256CBC = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 42}
	oidEncryptionAlgorithmAES128GCM = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 6}
	oidEncryptionAlgorithmAES128CBC = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 2}

	// 国密oid
	// SM4分组密码算法
	OidEncryptionAlgorithmSM4 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 104}
)

// 摘要算法oid
var (
	oidEncryptionAlgorithmSM3 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 401}
)

