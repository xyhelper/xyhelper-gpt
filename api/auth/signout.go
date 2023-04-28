package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func SignOut(r *ghttp.Request) {
	r.Cookie.Remove("access-token")
	r.Response.WriteJsonExit(g.Map{
		"url": "/login",
	})
}
