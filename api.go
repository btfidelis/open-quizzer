package main

import (
	"encoding/json"
	"github.com/btfidelis/quizzer/models"
	"github.com/btfidelis/quizzer/validation"
	"github.com/gocraft/web"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"net/http"
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

	encodedRet, err := json.Marshal(ret)

	if err != nil {
		log.Fatal("Unable to parse to Json ", err)
	}

	io.WriteString(rw, string(encodedRet))
}

func (a Api) GetQuestionList(rw web.ResponseWriter, req *web.Request) {
	quiz := models.Question{}
	a.Response(rw, req, quiz.ListAll(), 200)
}

func (a Api) CreateQuestion(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	quiz := models.Question{
		Type:           req.Form["type"][0],
		Question:       req.Form["question"][0],
		Answers:        req.Form["answers"],
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

	a.Response(rw, req, map[string]string{"status": "ok", "message": string(quiz.Id)}, http.StatusOK)
}

func (a Api) UpdateQuestion(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	quizId := req.PathParams["id"]
	quiz := models.Question{}
	err := quiz.Load(quizId)

	if err != nil {
		a.Response(rw, req, "", http.StatusNotFound)
	}

	quiz.Patch(req.Form)

	v := validation.ValidateQuiz(quiz)

	if !v.Passed {
		a.Response(rw, req, v.Errors, http.StatusBadRequest)
	}

	err = quiz.Update()
	if err != nil {
		a.Response(rw, req, "", http.StatusInternalServerError)
	}

	a.Response(rw, req, map[string]string{"status": "ok", "message": quiz.Id.String()}, http.StatusOK)
}

func (a Api) GetQuestion(rw web.ResponseWriter, req *web.Request) {
	quizId := req.PathParams["id"]
	quiz := models.Question{}
	err := quiz.Load(quizId)

	if err != nil {
		a.Response(rw, req, "", http.StatusInternalServerError)
	}

	a.Response(rw, req, quiz, http.StatusOK)
}

func (a Api) DeleteQuestion(rw web.ResponseWriter, req *web.Request) {
	quizId := req.PathParams["id"]
	quiz := models.Question{Id: bson.ObjectIdHex(quizId)}
	err := quiz.Delete()

	if err != nil {
		a.Response(rw, req, "", http.StatusInternalServerError)
	}

	a.Response(rw, req, quizId, http.StatusOK)
}

func (a Api) GetQuizList(rw web.ResponseWriter, req *web.Request) {
	quiz := models.Quiz{}
	a.Response(rw, req, quiz.ListAll(), 200)
}

func (a Api) GetQuiz(rw web.ResponseWriter, req *web.Request) {
	quizId := req.PathParams["id"]
}