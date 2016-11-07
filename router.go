package main

import (
	"github.com/gocraft/web"
)

func SetRoutes(r *web.Router) {
	r.Get("/question", (*Api).GetQuestionList).
		Get("/question/:id", (*Api).GetQuestion).
		Post("/question", (*Api).CreateQuestion).
		Patch("/question/:id", (*Api).UpdateQuestion).
		Delete("/question/:id", (*Api).DeleteQuestion)
	r.Get("/quiz", (*Api).GetQuizList)
}
