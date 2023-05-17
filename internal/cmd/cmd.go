package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"xyhelper-gpt/api/auth"
	"xyhelper-gpt/api/web"
	"xyhelper-gpt/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/hello", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.New(),
				)
			})
			s.SetServerRoot("resource/public/static")

			group := s.Group("/")
			group.GET("/", web.Chat)
			group.GET("/c/:ChatId", web.C)
			group.GET("/chat", web.Chat)
			group.GET("/login", web.Login)
			group.POST("/login", web.LoginPost)
			group.GET("/login_get", web.LoginPost)
			// group.GET("/api/auth/session", auth.Session)

			group.GET("/auth/logout", auth.LogOut)

			apiGroup := s.Group("/api")
			apiGroup.POST("/auth/logout", auth.LogOut)
			apiGroup.GET("/auth/session", auth.Session)
			PORT := g.Cfg().MustGetWithEnv(ctx, "PORT")
			g.Log().Debug(ctx, "PORT: "+PORT.String())
			s.SetPort(PORT.Int())

			s.Run()

			return nil
		},
	}
)
