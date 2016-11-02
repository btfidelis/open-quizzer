package main

import (
	"github.com/gocraft/web"
)

func SetRoutes(r *web.Router) {
	r.Get("/quiz", (*Api).GetQuizList).
 	  Post("/quiz", (*Api).CreateQuiz)

}