package open

const (
	GetAppAuthAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	GetAppPreAuthCodeUrl     = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	GetAuthInfoUrl           = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s"
	WxAppLoginUrl            = "https://api.weixin.qq.com/sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
	RefreshAccessTokenUrl    = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s"
)
