package mp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanming-open/wechat/utils"
	"io"
	"sort"
	"strings"
)

// @title encodingAESKey2AESKey
// @description
// @auth
// @param encodingKey string "加密串"
// @return data []byte "返回结果"
func encodingAESKey2AESKey(encodingKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingKey + "=")
	return data
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}
	return sha1Encode(b.String()) == signature
}

// 进行Sha1编码
func sha1Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func aesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}

func makeMsgSignature(token, timestamp, nonce, msg_encrypt string) string {
	sl := []string{token, timestamp, nonce, msg_encrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// 消息校验
func ValidateMsg(token, timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
	msgSignatureGen := makeMsgSignature(token, timestamp, nonce, msgEncrypt)
	return msgSignatureGen == msgSignatureIn
}

// 长连接换取短连接的方法,共用
func long2short(url string,accessToken string) string {
	posturl := fmt.Sprintf("%sshorturl?access_token=%s", wxApiHost,accessToken)
	var params = utils.KV{}
	params["action"] = "long2short"
	params["long_url"] = url
	buf, err := utils.DoPost(posturl, params)
	if err != nil {
		logger.Error(err.Error())
		return ""
	} else {
		var response = long2ShortResponse{}
		json.Unmarshal(buf, &response)
		if response.ErrCode > 0 {
			logger.Error(response.ErrMsg)
			return ""
		} else {
			return response.ShortUrl
		}
	}
}
