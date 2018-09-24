package controllers

import (
	"fmt"
	"log"

	"SimpleRestAPI/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCtr to export func
type UserCtr struct {
	Server   string
	Database string
}

var sess *mgo.Session
var db *mgo.Database

// COLLECTION DB USE
const (
	COLLECTION = "users"
)

// Connect use to connect MongoDB
func (m *UserCtr) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	sess = session
	db = session.DB(m.Database)
	fmt.Println(`
            MMMM      MMMM  MMMMMMMMM  MMMMM    MMM  MMMMMMMMM  MMMMMMMMM
            MMMMM    MMMMM  MMM   MMM  MMM MM   MMM  MMM    MM  MMM   MMM
            MMM MM  MM MMM  MMM   MMM  MMM  MM  MMM  MMM        MMM   MMM
            MMM  MMMM  MMM  MMM   MMM  MMM   MM MMM  MMM  MMMM  MMM   MMM
            MMM   MM   MMM  MMM   MMM  MMM    MMMMM  MMM    MM  MMM   MMM
            MMM        MMM  MMMMMMMMM  MMM     MMMM  MMMMMMMMM  MMMMMMMMM

CCCCCCCCC   CCCCCCCCC  CCCCC     CCC  CCCCC     CCC  CCCCCCCCC  CCCCCCCCC  CCCCCCCCC
CCC   CCC   CCC   CCC  CCC CC    CCC  CCC CC    CCC  CCC        CCC   CCC  CCCCCCCCC
CCC         CCC   CCC  CCC  CC   CCC  CCC  CC   CCC  CCCCCCCCC  CCC           CCC
CCC         CCC   CCC  CCC   CC  CCC  CCC   CC  CCC  CCC        CCC           CCC
CCC   CCC   CCC   CCC  CCC    CC CCC  CCC    CC CCC  CCC        CCC   CCC     CCC
CCCCCCCCC   CCCCCCCCC  CCC     CCCCC  CCC     CCCCC  CCCCCCCCC  CCCCCCCCC     CCC
	`)
}

// FindAll use to Find list of users
func (m *UserCtr) FindAll() ([]models.User, error) {
	var users []models.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// FindByID use to Find a user by its id
func (m *UserCtr) FindByID(id string) (models.User, error) {
	var user models.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserCtr) Insert(user models.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UserCtr) Delete(user models.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UserCtr) Update(user models.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}

// Close use to close mongodb connection
func (m *UserCtr) Close() {
	if sess != nil {
		sess.Close()
		sess = nil
		db = nil
		fmt.Println(`Connection Close`)
		return
	}
	fmt.Println(`Not Connect`)
}
