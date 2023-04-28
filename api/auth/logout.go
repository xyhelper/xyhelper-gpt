package auth

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func LogOut(r *ghttp.Request) {
	r.Cookie.Remove("access-token")
	// r.Response.WriteJsonExit(g.Map{
	// 	"url": "/login",
	// })
	r.Response.RedirectTo("/login")
}
