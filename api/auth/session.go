package auth

import (
	"xyhelper-gpt/config"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Session(r *ghttp.Request) {
	ctx := r.GetCtx()
	token := r.Cookie.Get("access-token")
	Expires := r.Cookie.Get("access-token-expires")
	c := g.Client()
	c.SetHeader("Authorization", "Bearer "+token.String())
	res, err := c.Post(ctx, config.API_PROXY+"/app/chatgpt/open/get_session", g.Map{
		"AccessToken": token.String(),
	})
	if err != nil {
		r.Response.Status = 500
		r.Response.WriteJsonExit(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	defer res.Close()
	if res.StatusCode != 200 {
		// r.Response.Status = res.StatusCode
		r.Response.WriteStatusExit(res.StatusCode)
	}
	resJson := gjson.New(res.ReadAllString())
	// g.Dump(resJson)
	if resJson.Get("code").Int() != 1000 {
		r.Cookie.Remove("access-token")
		r.Response.WriteStatusExit(401)
	}

	resJson.Set("data.user.name", Expires.String())
	r.Response.WriteJson(resJson.Get("data"))
}
