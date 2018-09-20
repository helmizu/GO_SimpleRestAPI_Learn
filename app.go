package main

import (
	"encoding/json"
	"log"
	"net/http"

	conf "SimpleRestAPI/config"
	ctrllr "SimpleRestAPI/controllers"
	model "SimpleRestAPI/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = conf.Config{}
var ctr = ctrllr.UserCtr{}

// AllUsersEndPoint use to GET all users
func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := ctr.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

// FindUserEndpoint use to GET a users by its ID
func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := ctr.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// CreateUserEndPoint use to POST a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := ctr.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

// UpdateUserEndPoint use to PUT update an existing user
func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var user model.User
	user.ID = bson.ObjectIdHex(params["id"])
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ctr.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserEndPoint use to DELETE an existing user
func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ctr.Delete(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Init()
	ctr.Server = config.Server
	ctr.Database = config.Database
	ctr.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users/{id}", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
