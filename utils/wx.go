package utils

import (
	"encoding/base64"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"strings"
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

	result = ConvertContentToXml(string(result))

	return gxml.Decode(result)
}

func ConvertContentToXml(content string) []byte {
	var (
		tmp []string
	)

	tmp = strings.Split(content, "<xml>")
	tmp = strings.Split(tmp[1], "</xml>")

	return []byte("<xml>" + tmp[0] + "</xml>")
}

