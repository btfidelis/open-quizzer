package main

import (
	"github.com/gocraft/web"
	"net/http"
	"io"
	"encoding/json"
	"log"
	"github.com/btfidelis/quizzer/models"
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
