package RSA

import (
	"campushelphub/internal/config"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
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

// 解析公钥
func parsePublicKeyFromPEM(publicKeyStr string) (*rsa.PublicKey, error) {
	// 1. 空值校验
	if publicKeyStr == "" {
		return nil, errors.New("RSA公钥字符串不能为空")
	}

	// 2. 解码PEM格式数据（提取公钥的DER编码内容）
	block, rest := pem.Decode([]byte(publicKeyStr))
	if block == nil || len(rest) > 0 {
		return nil, errors.New("解析PEM格式公钥失败:无效的PEM数据(可能包含多余内容)")
	}

	// 3. 校验PEM块类型，仅处理合法的公钥类型
	switch block.Type {
	case "RSA PUBLIC KEY": // PKCS1格式
		// 解析PKCS1格式公钥
		pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析PKCS1格式RSA公钥失败:%w", err)
		}
		return pubKey, nil

	case "PUBLIC KEY": // PKIX/X.509格式（最常用）
		// 解析PKIX格式公钥
		pubKeyIface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析PKIX格式RSA公钥失败:%w", err)
		}
		// 类型断言为RSA公钥（排除ECC等其他类型公钥）
		rsaPubKey, ok := pubKeyIface.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("解析出的公钥不是RSA类型(可能是ECC等其他算法)")
		}
		return rsaPubKey, nil

	default:
		return nil, fmt.Errorf("不支持的PEM块类型:%s(需为RSA PUBLIC KEY或PUBLIC KEY)", block.Type)
	}
}

// 构造函数
func NewRSA(config *config.Config) *RSA {
	privateKey, err := parsePrivateKeyFromPEM(config.RSA.PrivateKey)
	publicKey, err := parsePublicKeyFromPEM(config.RSA.PublicKey)
	if err != nil {
		return nil
	}
	return &RSA{
		privateKey: privateKey,
		publicKey:  publicKey,
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
