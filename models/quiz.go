package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Quiz struct {
	Id   bson.ObjectId  `bson:"_id,omitempty"`
	Name string         `bson:"name"`
	Slug string         `bson:"slug"`
	Questions []bson.ObjectId   `bson:"questions"`
}

func (q Quiz) getCollection() string {
	return "quizzes"
}

func (q Quiz) ListAll() []Quiz {
	var quizzes []Quiz
	getCollectionFromModel(q, func(c *mgo.Collection) {
		c.Find(nil).All(quizzes)
	})

	return quizzes
}


func (q *Quiz) Load(id string) error {
	var err error
	getCollectionFromModel(q, func(c *mgo.Collection) {
		c.FindId(bson.ObjectIdHex(id)).One(q)
	})

	return err
}

