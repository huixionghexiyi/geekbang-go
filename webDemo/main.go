package main

import (
	"geekbang-go/webDemo/handler"
	server2 "geekbang-go/webDemo/server"
)

func main() {
	server := server2.NewHttpServer("web demo")
	server.Route("GET", "/hello", handler.Hello)

	server.Start(":8080")
}
