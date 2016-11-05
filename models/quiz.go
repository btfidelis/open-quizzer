package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (q *Quiz) Patch(req map[string][]string) {
	quiz := Quiz{
		Type: req["type"][0],
		Question: req["question"][0],
		Answers:  req["answers"],
		CorrectAnswers: req["correct_answers"],
	}

	if q.Type != quiz.Type {
		q.Type = quiz.Type
	}

	if q.Question != quiz.Question {
		q.Question = quiz.Question
	}

	if !StringSliceEquals(q.Answers, quiz.Answers) {
		q.Answers = quiz.Answers
	}

	if !StringSliceEquals(q.CorrectAnswers, quiz.CorrectAnswers) {
		q.CorrectAnswers = quiz.CorrectAnswers
	}


}

func (q *Quiz) Update() error {
	var err error
	getCollectionFromModel(q, func(c *mgo.Collection) {
		err = c.UpdateId(q.Id, q)
	})

	return err
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
