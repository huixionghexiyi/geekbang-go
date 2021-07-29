package context

import (
	"encoding/json"
	"io"
	"net/http"
)

// Context 上下文
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

func (c *Context) ReadJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, req)
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) SystemErrorJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func (c *Context) SuccessJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}
