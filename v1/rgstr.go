package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Rgstr(r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"success": true,
	})
}
