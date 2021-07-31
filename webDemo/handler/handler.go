package handler

import (
	"geekbang-go/webDemo/context"
	"net/http"
)

type Routable interface {
	Route(method, pattern string, handlerFunc func(c *context.Context))
}

type Handler interface {
	http.Handler
	Routable
}

// HandlerBasedOnMap 实现了 Handler 接口，通过注册 handler 的请求，都需要通过 ServeHTTP 方法
type HandlerBasedOnMap struct {
	// key 应该是 method + url
	Handlers map[string]func(c *context.Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key(request.Method, request.URL.Path)
	if handler, ok := h.Handlers[key]; ok {
		handler(context.NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (h *HandlerBasedOnMap) Route(method, pattern string, handlerFunc func(c *context.Context)) {
	key := h.Key(method, pattern)
	h.Handlers[key] = handlerFunc
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		Handlers: make(map[string]func(c *context.Context)),
	}
}

func (h *HandlerBasedOnMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}
