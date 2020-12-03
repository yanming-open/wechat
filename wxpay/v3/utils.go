package v3

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	mrand "math/rand"
)

// 加密
func rsa256Encrypt(ciphertext, keyPath string) string {
	privKey, _ := ioutil.ReadFile(keyPath)
	KeyBlock, _ := pem.Decode(privKey)

	random := rand.Reader
	msg := []byte(ciphertext)
	hash := sha256.New()
	hash.Write(msg)
	privateKey, err := x509.ParsePKCS8PrivateKey(KeyBlock.Bytes)
	if err != nil {
		logger.Error(err.Error())
	}
	sign, err := rsa.SignPKCS1v15(random, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash.Sum(nil))
	if err != nil {
		logger.Error(err.Error())
	}
	return base64.StdEncoding.EncodeToString(sign)
}

//
func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[mrand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
