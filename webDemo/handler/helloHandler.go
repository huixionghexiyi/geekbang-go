package handler

import (
	"geekbang-go/webDemo/context"
	"time"
)


func Hello(c *context.Context){
	// TODO 增加header头：text/html
	c.SuccessJson("hello world")
	time.Sleep(time.Second)
}