package main

import (
	"github.com/gocraft/web"
)

func SetMiddleware(r *web.Router) {
	r.Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware)
}
