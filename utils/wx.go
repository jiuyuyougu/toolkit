package utils

import (
	"encoding/base64"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"log"
)

func EncodingAESKey2AESKey(encodingKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingKey + "=")
	return data
}

func DecryptRequest(msg, key string) (g.Map, error) {
	byteKey := EncodingAESKey2AESKey(key)

	content, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	result, err := AesDecrypt(content, byteKey)
	if err != nil {
		return nil, err
	}

	log.Println(result)

	return gxml.Decode(result)
}
