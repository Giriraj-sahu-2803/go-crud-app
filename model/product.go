package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Price int           `json:"price" bson:"price"`
}
