package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	API_PROXY = "http://127.0.0.1:8001"
)

func init() {
	ctx := gctx.GetInitCtx()
	apiProxy := g.Cfg().MustGetWithEnv(ctx, "API_PROXY").String()
	if apiProxy != "" {
		API_PROXY = apiProxy
	}
	g.Log().Debug(ctx, "API_PROXY: "+API_PROXY)
}
