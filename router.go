package main

import (
	"github.com/gocraft/web"
)

func SetRoutes(r *web.Router) {
	r.Get("/quiz", (*Api).GetQuizList).
	  Get("/quiz/:id", (*Api).GetQuiz).
 	  Post("/quiz", (*Api).CreateQuiz).
	  Patch("/quiz/:id", (*Api).UpdateQuiz)

}