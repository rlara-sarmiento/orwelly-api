package router

import "net/http"

type Route struct {
	Path        string
	Method      string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}
