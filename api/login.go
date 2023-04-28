package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// LoginToken
func LoginToken(r *ghttp.Request) {
	req := r.GetMapStrStr()
	if req["action"] == "token" && req["access_token"] != "" {
		r.Cookie.Set("access-token", req["access_token"])
		r.Response.WriteJsonExit(g.Map{
			"code": 0,
			"url":  "/",
		})
	} else {
		r.Response.WriteJsonExit(g.Map{
			"code": 500,
			"msg":  "invalid token",
		})
	}
}
