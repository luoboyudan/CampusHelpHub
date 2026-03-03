package RSA

import (
	"campushelphub/internal/config"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type RSA struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func parsePrivateKeyFromPEM(privateKeyStr string) (*rsa.PrivateKey, error) {
	// 1. 空值校验
	if privateKeyStr == "" {
		return nil, errors.New("RSA私钥字符串不能为空")
	}

	// 2. 解码PEM格式数据
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("解析PEM格式私钥失败:无效的PEM数据")
	}

	// 3. 校验PEM块类型（确保是RSA私钥）
	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("PEM块类型错误:需要RSA PRIVATE KEY")
	}

	// 4. 解析ASN.1 DER格式为RSA私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// 兼容PKCS8格式的私钥（部分工具生成的私钥是PKCS8）
		pkcs8Key, pkcs8Err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if pkcs8Err != nil {
			return nil, errors.Join(errors.New("解析RSA私钥失败"), err, pkcs8Err)
		}
		// 类型断言为RSA私钥
		rsaPkcs8Key, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("PKCS8格式私钥不是RSA类型")
		}
		privateKey = rsaPkcs8Key
	}

	return privateKey, nil
}

// 构造函数
func NewRSA(config *config.Config) *RSA {
	privateKey, err := parsePrivateKeyFromPEM(config.RSA.PrivateKey)
	if err != nil {
		return nil
	}
	return &RSA{
		privateKey: privateKey,
	}
}

// 解密
func (r *RSA) Decrypt(encrypted []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), nil, r.privateKey, encrypted, nil)
}

// 获取公钥
func (r *RSA) GetPublicKey() string {
	if r.publicKey == nil {
		return ""
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(r.publicKey)
	if err != nil {
		return ""
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pubKeyPEM)
}
