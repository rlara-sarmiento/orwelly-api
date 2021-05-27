package router

import (
	"github.com/gorilla/mux"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/handler"
)

var routes = []Route{
	{Path: "/hello", Method: "GET", HandlerFunc: handler.Hello},
	{Path: "/quotes", Method: "POST", HandlerFunc: handler.CreateQuote},
}

func Get() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
	}
	r.HandleFunc("/", handler.NotFound).Methods()
	return r
}
