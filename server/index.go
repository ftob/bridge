package server

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type indexHandler struct {}

func (ih *indexHandler) handler(ctx *routing.Context) error {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/json")
	ctx.SetBody([]byte("{}"))
	return nil
}