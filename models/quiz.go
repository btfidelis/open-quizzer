package models

import (
	"gopkg.in/mgo.v2"
)

type Quiz struct {
	Question string  `bson:"question"`
	Answers []string `bson:"answers"`
}

func (q Quiz) getCollection() string {
	return "quizzes"
}

func (q Quiz) ListAll() []Quiz {
	var quizList []Quiz
	getCollectionFromModel(Quiz{}, func(c *mgo.Collection) {
		c.Find(nil).All(&quizList)
	})

	return quizList
}
