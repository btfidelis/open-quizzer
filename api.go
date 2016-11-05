package main

import (
	"github.com/gocraft/web"
	"net/http"
	"io"
	"encoding/json"
	"log"
	"github.com/btfidelis/quizzer/models"
	"github.com/btfidelis/quizzer/validation"
)

type Api struct {
	ResponseDefaultEncoding string
}


func (a Api) Response(rw web.ResponseWriter, req *web.Request, ret interface{}, httpCode int) {
	rw.Header().Set("Content-Type", a.ResponseDefaultEncoding)

	if httpCode >= 400 {
		http.Error(rw, http.StatusText(httpCode), httpCode)
		return
	} else {
		rw.WriteHeader(httpCode)
	}

	encodedRet, err :=  json.Marshal(ret)

	if (err != nil) {
		log.Fatal("Unable to parse to Json ", err)
	}

	io.WriteString(rw, string(encodedRet))
}

func (a Api) GetQuizList(rw web.ResponseWriter, req *web.Request) {
	quiz := models.Quiz{}
	a.Response(rw, req, quiz.ListAll(), 200)
}

func (a Api) CreateQuiz(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	quiz := models.Quiz{
		Type: req.Form["type"][0],
		Question: req.Form["question"][0],
		Answers:  req.Form["answers"],
		CorrectAnswers: req.Form["correct_answers"],
	}

	v := validation.ValidateQuiz(quiz)

	if !v.Passed {
		a.Response(rw, req, v.Errors, http.StatusBadRequest)
	}

	err := quiz.Save()
	if err != nil {
		a.Response(rw, req, "", http.StatusInternalServerError)
	}

	a.Response(rw, req, map[string]string { "status": "ok", "message": string(quiz.Id) }, http.StatusOK)
}

func (a Api) UpdateQuiz(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	quizId := req.PathParams["id"]
	quiz := new(models.Quiz)
	err := quiz.Load(quizId)

	if err != nil {
		a.Response(rw, req, "", http.StatusNotFound)
	}

	quiz.Patch(req.Form)

	v := validation.ValidateQuiz(quiz)

	if !v.Passed {
		a.Response(rw, req, v.Errors, http.StatusBadRequest)
	}

	err = quiz.Save()
	if err != nil {
		a.Response(rw, req, "", http.StatusInternalServerError)
	}

	a.Response(rw, req, map[string]string {"status": "ok", "message": quiz.Id.String()} , http.StatusOK)
}
