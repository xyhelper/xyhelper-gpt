package backendapi

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"pandora-go/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	u, _ := url.Parse(config.API_PROXY)
	s.BindHandler("/backend-api/*", func(r *ghttp.Request) {
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
			writer.WriteHeader(http.StatusBadGateway)
		}
		proxy.ServeHTTP(r.Response.Writer.RawWriter(), r.Request)
	})

}
