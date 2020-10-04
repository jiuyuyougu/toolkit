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

// 获取授权信息
func GetAuthInfo(accessToken, authCode, appID string) (g.Map, error) {
	reqMap := g.Map{
		"component_appid":    appID,
		"authorization_code": authCode,
	}

	reqBody, _ := gjson.Encode(reqMap)

	rsp, err := http.Post(
		fmt.Sprintf(GetAuthInfoUrl, accessToken),
		"application/json",
		bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	rspData, err := gjson.Decode(rspBody)
	if err != nil {
		return nil, err
	}

	rspMap := gconv.Map(rspData)

	return gconv.Map(rspMap["authorization_info"]), nil
}

// 刷新token
func RefreshAccessToken(accessToken, AppID, AuthAppID, refreshToken string) (g.Map, error) {
	url := fmt.Sprintf(
		RefreshAccessTokenUrl, accessToken)

	reqMap := g.Map{
		"component_appid":          AppID,
		"authorizer_appid":         AuthAppID,
		"authorizer_refresh_token": refreshToken,
	}

	reqBody, err := gjson.Encode(reqMap)
	if err != nil {
		return nil, err
	}

	rsp, err := http.Post(
		url,
		"application/json",
		bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	rspData, err := gjson.Decode(rspBody)
	if err != nil {
		return nil, err
	}

	return gconv.Map(rspData), nil
}

// 登录
func WxAppLogin(appID, code, comAppID, accessToken string) (g.Map, error) {
	url := fmt.Sprintf(
		WxAppLoginUrl,
		appID, code, comAppID, accessToken)

	rsp, err := http.Post(
		url,
		"application/json",
		nil)
	if err != nil {
		return nil, err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	rspData, err := gjson.Decode(rspBody)
	if err != nil {
		return nil, err
	}

	return gconv.Map(rspData), nil
}
