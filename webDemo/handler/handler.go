package handler

import (
	"geekbang-go/webDemo/context"
	"net/http"
)

// BasedHandlerOnMap 实现了 Handler 接口，通过注册 handler 的请求，都需要通过 ServeHTTP 方法
type BasedHandlerOnMap struct {
	// key 应该是 method + url
	Handlers map[string]func(c *context.Context)
}

func (h *BasedHandlerOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key(request.Method, request.URL.Path)
	if handler, ok := h.Handlers[key]; ok {
		handler(context.NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (h *BasedHandlerOnMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}
