package server

import (
	"fmt"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// server and routing
type Server struct {
	router *routing.Router
}

// New server and make routing map
func New() *Server {
	router := routing.New()
	// not found page handler
	nfh := &notFoundHandler{}
	router.NotFound(nfh.handler)

	// index handler
	ih := &indexHandler{}
	router.Get("/", ih.handler)

	iat := &iotaHandler{}
	router.Get("/iota", iat.handler)

	srv := &Server{
		router: router,
	}
	return srv
}

// Listen request
func (s *Server) ListenAndServe(port uint) error {
	addr := fmt.Sprintf(":%d", port)

	return fasthttp.ListenAndServe(addr, s.router.HandleRequest)
}