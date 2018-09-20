package controllers

import (
	"log"

	. "SimpleRestAPI/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserCtr struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UserCtr) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of users
func (m *UserCtr) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its id
func (m *UserCtr) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserCtr) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UserCtr) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UserCtr) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
