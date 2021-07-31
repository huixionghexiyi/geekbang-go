package server

import (
	"geekbang-go/webDemo/context"
	"geekbang-go/webDemo/handler"
)

// Server 的抽象
type Server interface {
	handler.Routable
	// Start 启动一个Server
	Start(address string) error
}

// handlerFunc 处理方法的抽象
type handlerFunc func(c *context.Context)
