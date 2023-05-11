package backendapi

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"xyhelper-gpt/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	u, _ := url.Parse(config.API_PROXY)
	s.BindHandler("/api/*", func(r *ghttp.Request) {
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
			writer.WriteHeader(http.StatusBadGateway)
		}
		req := r.Request
		// 替换 r.URL.Path 中的 /api 为 /backend-api
		// 例如：/api/auth/session -> /backend-api/auth/session
		req.URL.Path = "/backend-api" + req.URL.Path[4:]
		req.Host = u.Host // 替换请求头中的 Host 为 API_PROXY 的 Host
		g.Log().Debug(r.GetCtx(), "proxy to "+u.String()+req.URL.Path)

		proxy.ServeHTTP(r.Response.Writer.RawWriter(), r.Request)
	})
	s.BindHandler("/api/conversation_limit", func(r *ghttp.Request) {
		res := `
		{
			"message_cap": 25,
			"message_cap_window": 180,
			"message_disclaimer": {
				"model-switcher": "您已达到 GPT-4 上限，这为所有 ChatGPT Plus 用户提供了尝试该模型的机会。\n\n请稍后再试。",
				"textarea": "目前 GPT-4 限制 每 3 小时最多处理 25 条消息。"
			}
		}
	`
		r.Response.WriteJsonExit(res)
	})

}
