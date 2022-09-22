package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// RSA 加密
func RsaEncrypt(plaintext string, pubKey *rsa.PublicKey) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(plaintext))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Rsa 解密
func RsaDecrypt(ciphertext string, prikey *rsa.PrivateKey) (string, error) {
	cipherbytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptPKCS1v15(nil, prikey, cipherbytes)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// 载入私钥从Base64
func RsaLoadPrivateKeyBase64(base64key string) (*rsa.PrivateKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(keybytes)
	if err == nil {
		return privatekey, err
	}

	privatekey2, err := x509.ParsePKCS8PrivateKey(keybytes)
	if err != nil {
		return nil, err
	}
	return privatekey2.(*rsa.PrivateKey), nil
}

// 载入公钥从Base64
func RsaLoadPublicKeyBase64(base64key string) (*rsa.PublicKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	pubkeyinterface, err := x509.ParsePKIXPublicKey(keybytes)
	if err != nil {
		return nil, err
	}
	publickey := pubkeyinterface.(*rsa.PublicKey)

	// publickey, err := x509.ParsePKCS1PublicKey(keybytes)
	// if err != nil {
	// 	return nil, err
	// }

	return publickey, nil
}
