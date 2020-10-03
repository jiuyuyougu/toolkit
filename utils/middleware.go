package utils

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func MiddlewareCors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MiddlewareSetUserID(r *ghttp.Request) {
	rsp := NewResp(r)
	auth := gconv.Map(r.GetCtxVar("jy-auth"))

	if auth == nil {
		rsp.UNAUTHORIZED("授权信息不存在！请重新登录！")
		return
	}

	userInfo := gconv.Map(auth["user_info"])
	r.SetCtxVar("user_info", userInfo)
	r.Middleware.Next()
}
