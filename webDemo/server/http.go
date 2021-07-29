package server

import (
	"geekbang-go/webDemo/context"
	"net/http"
)

// sdkHttpServer 实现Server
type sdkHttpServer struct {
	Name string

}

// Route 路由
func (s *sdkHttpServer) Route(pattern string, handlerFunc handlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(w, r)
		handlerFunc(ctx)
	})
}

// Start 启动服务
func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

// NewHttpServer 创建一个服务
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
