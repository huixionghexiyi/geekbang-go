package server

import "geekbang-go/webDemo/context"

// Server 的抽象
type Server interface {
	// Route 设定一个路由，命中该陆游执行handlerFunc 代码
	Route(method, pattern string, handlerFunc handlerFunc)

	// Start 启动一个Server
	Start(address string) error
}

// handlerFunc 处理方法的抽象
type handlerFunc func(c *context.Context)
