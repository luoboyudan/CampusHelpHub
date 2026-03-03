package RSA

import (
	"campushelphub/internal/config"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

type RSA struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func parsePrivateKeyFromPEM(privateKeyStr string) (*rsa.PrivateKey, error) {
	if privateKeyStr == "" {
		return nil, errors.New("RSA私钥字符串不能为空")
	}

	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("解析PEM格式私钥失败:无效的PEM数据")
	}

	switch block.Type {
	case "RSA PRIVATE KEY": // PKCS1格式
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析PKCS1格式私钥失败: %w", err)
		}
		return privateKey, nil

	case "PRIVATE KEY": // PKCS8格式
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析PKCS8格式私钥失败: %w", err)
		}
		rsaKey, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("PKCS8格式私钥不是RSA算法")
		}
		return rsaKey, nil

	default: // 不支持的格式
		return nil, fmt.Errorf("不支持的PEM块类型: %s,仅支持RSA PRIVATE KEY(PKCS1)或PRIVATE KEY(PKCS8)", block.Type)
	}
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
	if err != nil {
		panic(err)
	}
	publicKey, err := parsePublicKeyFromPEM(config.RSA.PublicKey)
	if err != nil {
		panic(err)
	}
	return &RSA{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// 解密
func (r *RSA) Decrypt(encrypted string) ([]byte, error) {
	decryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, fmt.Errorf("解码base64字符串失败:%w", err)
	}

	return rsa.DecryptOAEP(sha256.New(), nil, r.privateKey, decryptedBytes, nil)
}

// 获取公钥
func (r *RSA) GetPublicKey() (string, error) {
	if r.publicKey == nil {
		return "", errors.New("RSA公钥不存在")
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(r.publicKey)
	if err != nil {
		return "", fmt.Errorf("序列化PKIX公钥失败:%w", err)
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pubKeyPEM), nil
}
