package auth

import (
	"xyhelper-gpt/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Session(r *ghttp.Request) {
	ctx := r.GetCtx()
	token := r.Cookie.Get("access-token")
	c := g.Client()
	c.SetHeader("Authorization", "Bearer "+token.String())
	res, err := c.Get(ctx, config.API_PROXY+"/api/auth/session")
	if err != nil {
		r.Response.Status = 500
		r.Response.WriteJsonExit(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	defer res.Close()
	if res.StatusCode != 200 {
		r.Response.Status = res.StatusCode
	}

	r.Response.WriteJson(res.ReadAllString())
}
