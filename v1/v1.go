package v1

import "github.com/gogf/gf/v2/frame/g"

func init() {
	s := g.Server()
	v1group := s.Group("/v1")
	v1group.ALL("/rgstr", Rgstr)
	v1group.ALL("/initialize", Initialize)

}
