package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gocraft/web"
)

const (
	QUIZ_MULTIPLE_ALTERNATIVE = "M"
	QUIZ_NORMAL = "N"
)

type Quiz struct {
	Id               bson.ObjectId  `bson:"_id,omitempty"`
	Type             string         `bson:"type"`
	Question         string         `bson:"question"`
	Answers          []string       `bson:"answers"`
	CorrectAnswers   []string       `bson:"correct_answers"`
}

func (q Quiz) getCollection() string {
	return "quizzes"
}

func (q *Quiz) Patch(req *web.Request) {
	quiz := Quiz{
		Type: req.Form["type"][0],
		Question: req.Form["question"][0],
		Answers:  req.Form["answers"],
		CorrectAnswers: req.Form["correct_answers"],
	}

	if q.Type != quiz.Type {
		q.Type = quiz.Type
	}

	if q.Question != quiz.Question {
		q.Type = quiz.Type
	}

	if q.Answers != quiz.Answers {
		q.Type = quiz.Type
	}

	if q.CorrectAnswers != quiz.CorrectAnswers {
		q.Type = quiz.Type
	}
}

func (q *Quiz) ListAll() []Quiz {
	var quizList []Quiz
	getCollectionFromModel(q, func(c *mgo.Collection) {
		c.Find(nil).All(&quizList)
	})

	return quizList
}

func (q *Quiz) Load(id string) error {
	var err error
	getCollectionFromModel(q, func(c *mgo.Collection) {
		err = c.FindId(bson.ObjectIdHex(id)).One(q)
	})

	return err
}

func (q *Quiz) Save() error {
	var err error
	q.Id = bson.NewObjectId()

	getCollectionFromModel(q, func(c *mgo.Collection) {
		err = c.Insert(q)
	})

	return err
}
