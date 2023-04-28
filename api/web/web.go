package web

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Chat(r *ghttp.Request) {
	if r.Cookie.Get("access-token").String() == "" {
		g.Log().Debug(r.GetCtx(), "redirect to login")
		r.Response.RedirectTo("/login")
	}
	r.Response.WriteTpl("chat.html", g.Map{
		"pandora_sentry": false,
		"api_prefix":     "https://ai.fakeopen.com",
		"props":          `{"buildId": "tTShkecJDS0nIc9faO2vC", "gssp": true, "isFallback": false, "page": "/chat/[[...chatId]]", "props": {"__N_SSP": true, "pageProps": {"geoOk": true, "isUserInCanPayGroup": true, "serviceAnnouncement": {"paid": {}, "public": {}}, "serviceStatus": {}, "user": {"email": "a269@xyhelper.cn", "groups": [], "id": "user-VbBwHhwqj4rez2SD2D2QmZkM", "image": null, "name": "a269@xyhelper.cn", "picture": null}, "userCountry": "US"}}, "query": {}, "scriptLoader": []}`,
	})
}

func Login(r *ghttp.Request) {
	g.Log().Debug(r.GetCtx(), "login")
	r.Response.WriteTpl("login.html", g.Map{})
}
