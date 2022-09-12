package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"strconv"

	"my_project/model"
	"net/http"
)

type UserCtrl struct {
	users model.UserModel
	//DB    *sql.DB
}

//type User struct {
//	Id   int    `json:"id"`
//	Name string `json:"name"`
//	Sale int    `json:"sale"`
//}
var usr = &model.UserModel{}

func NewUserCtrl(*sql.DB) *UserCtrl {

	return &UserCtrl{}
}

func (usr *UserCtrl) Getusers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	users, err := usr.users.Getusers()
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	//p,err := usr.users.Getusers()
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	json.NewEncoder(res).Encode(&users)
}

func (usr *UserCtrl) GetSingleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var id = params["id"]
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}

	p, err := usr.users.GetSingleUser(s)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}

	json.NewEncoder(res).Encode(&p)

}
func (usr *UserCtrl) CreateUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//
	//var param model.User
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {

		log.Printf("Got error in mysql connector: %s", err)
		return
	}

	m, err := usr.users.CreateUser(user.Name, user.Sale)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)

		return
	}
	json.NewEncoder(res).Encode(&m)
}
func (usr *UserCtrl) UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	users, err := usr.users.UpdateUser(user.ID, user.Name, user.Sale)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}

	json.NewEncoder(res).Encode(&users)

}
func (usr *UserCtrl) DeleteUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//var user model.User
	params := mux.Vars(req)
	id := params["id"]
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	p, err := usr.users.DeleteUser(s)
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	json.NewEncoder(res).Encode(p)
}
