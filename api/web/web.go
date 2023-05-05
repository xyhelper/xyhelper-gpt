package web

import (
	"xyhelper-gpt/config"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
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
	r.Response.WriteTpl("login.html", g.Map{})
}

func LoginPost(r *ghttp.Request) {
	ctx := r.GetCtx()
	cli := g.Client()
	res, err := cli.Post(ctx, config.API_PROXY+"/app/chatgpt/open/check_user", g.Map{
		"AccessToken": r.Get("password"),
	})
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteTpl("login.html", g.Map{
			"Error": err.Error(),
		})
		return
	}
	defer res.Close()
	if res.StatusCode != 200 {
		g.Log().Error(ctx, res.StatusCode)
		r.Response.WriteTpl("login.html", g.Map{
			"Error": "服务暂时不可用" + res.Status,
		})
		return
	}
	resString := res.ReadAllString()
	g.Log().Debug(ctx, resString)
	resJson := gjson.New(resString)
	if resJson.Get("code").Int() != 1000 {
		r.Response.WriteTpl("login.html", g.Map{
			"Error": resJson.Get("message").String(),
		})
		return
	}
	if resJson.Get("data.bindMode").String() != "personal" {
		r.Response.WriteTpl("login.html", g.Map{
			"Error": "共享池用户无法使用",
		})
		return
	}
	ExpireTime := gtime.New(resJson.Get("data.ExpireTime").String())
	g.Log().Debug(ctx, ExpireTime, gtime.Now())
	if ExpireTime.Before(gtime.Now()) {
		r.Response.WriteTpl("login.html", g.Map{
			"Error": "AccessToken已过期",
		})
		return
	}
	r.Cookie.Set("access-token", r.Get("password").String())
	r.Response.RedirectTo("/")

}
