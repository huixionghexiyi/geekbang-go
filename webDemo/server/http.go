package server

import (
	"geekbang-go/webDemo/context"
	"geekbang-go/webDemo/handler"
	"net/http"
)

// sdkHttpServer 实现Server
type sdkHttpServer struct {
	Name    string
	Handler *handler.BasedHandlerOnMap
}

// Route 路由
func (s *sdkHttpServer) Route(method, pattern string, handlerFunc handlerFunc) {
	key := s.Handler.Key(method, pattern)
	s.Handler.Handlers[key] = handlerFunc
	//http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
	//	ctx := context.NewContext(w, r)
	//	handlerFunc(ctx)
	//})
}

// Start 启动服务
func (s *sdkHttpServer) Start(address string) error {
	// 启动服务
	return http.ListenAndServe(address, s.Handler)
}

// NewHttpServer 创建一个服务
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name:    name,
		Handler: &handler.BasedHandlerOnMap{
			Handlers: make(map[string]func(c *context.Context)),
		},
	}
}
