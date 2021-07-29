package handler

import (
	"encoding/json"
	"fmt"
	"geekbang-go/webDemo/context"
	"io"
)

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

func SignUp(c *context.Context) {
	req := &signUpReq{}
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		fmt.Fprintf(c.W, "read body failed: %v", err)
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(c.W, "deserinlized failed: %v", err)
		return
	}
	return
}
