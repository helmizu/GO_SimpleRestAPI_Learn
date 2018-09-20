package models

import "gopkg.in/mgo.v2/bson"

//User model
type User struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Age   int           `bson:"age" json:"age"`
	Email string        `bson:"email" json:"email"`
}
