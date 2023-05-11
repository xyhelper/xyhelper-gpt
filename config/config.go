package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	API_PROXY = "https://personalchat.xyhelper.cn"
	PORT      = 8800
)

func init() {
	ctx := gctx.GetInitCtx()
	apiProxy := g.Cfg().MustGetWithEnv(ctx, "API_PROXY").String()
	if apiProxy != "" {
		API_PROXY = apiProxy
	}
	g.Log().Debug(ctx, "API_PROXY: ", API_PROXY)
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port != 0 {
		PORT = port
	}
	g.Log().Debug(ctx, "PORT: ", PORT)
}
