package open

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/jiuyuyougu/toolkit/utils"
	"io/ioutil"
	"net/http"
)

// 获取验证票据
func GetAppComponentVerifyTicket(msg, key string) (string, error) {
	if msg != "" {
		res, err := utils.DecryptRequest(msg, key)
		if err != nil {
			return "", err
		}

		return gconv.String(
			gconv.Map(res["xml"])["ComponentVerifyTicket"],
		), nil

	} else {
		return "", errors.New("消息为空！")
	}
}

// 获取令牌
func GetAppAuthAccessToken(appID, appSecret, ticket string) (string, error) {
	reqMap := make(g.Map)

	reqMap["component_appid"] = appID
	reqMap["component_appsecret"] = appSecret
	reqMap["component_verify_ticket"] = ticket

	reqBody, _ := gjson.Encode(reqMap)

	rsp, err := http.Post(
		GetAppAuthAccessTokenUrl,
		"application/json",
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	rspData, err := gjson.Decode(rspBody)
	if err != nil {
		return "", err
	}

	rspMap := gconv.Map(rspData)

	return gconv.String(rspMap["component_access_token"]), nil
}

// 获取预授权码
func GetAppPreAuthCode(token, appID string) (string, error) {
	reqMap := make(g.Map)

	reqMap["component_appid"] = appID

	reqBody, _ := gjson.Encode(reqMap)

	rsp, err := http.Post(
		fmt.Sprintf(GetAppPreAuthCodeUrl, token),
		"application/json",
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	rspData, err := gjson.Decode(rspBody)
	if err != nil {
		return "", err
	}

	rspMap := gconv.Map(rspData)

	return gconv.String(rspMap["pre_auth_code"]), nil
}
