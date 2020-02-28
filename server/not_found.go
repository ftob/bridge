package server

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)
//
type notFoundHandler struct {}

func (nfh *notFoundHandler) handler(ctx *routing.Context) error {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	ctx.SetContentType("text/json")
	ctx.SetBody([]byte(`{"message": "Page not found"}`))
	return nil
}