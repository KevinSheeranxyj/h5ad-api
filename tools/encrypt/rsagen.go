// Package encrypt rsa秘钥对生成
package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GenRsaKey 生成Ras秘钥对
func GenRsaKey(bits int, path string) error {
	// 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "private.pem",
		Bytes: derStream,
	}
	file, err := os.Create(path + "private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 公钥
	publicKey := &privateKey.PublicKey // 私有key基础上生成公有key
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "public.pem",
		Bytes: derPkix,
	}
	file, err = os.Create(path + "public.pem")

	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
