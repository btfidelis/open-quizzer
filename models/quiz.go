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

func (q *Quiz) ListAll() []Quiz {
	var quizList []Quiz
	getCollectionFromModel(q, func(c *mgo.Collection) {
		c.Find(nil).All(&quizList)
	})

	return quizList
}

func (q *Quiz) Replace(where bson.M) {

}

func (q *Quiz) Save() error {
	var err error
	getCollectionFromModel(q, func(c *mgo.Collection) {
		err = c.Insert(q)
	})

	return err
}
