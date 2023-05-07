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
		g.Log().Debug(r.GetCtx(), "proxy to "+u.String()+req.URL.Path)
		req.URL.Path = "/backend-api" + req.URL.Path[4:]
		proxy.ServeHTTP(r.Response.Writer.RawWriter(), r.Request)
	})

}
