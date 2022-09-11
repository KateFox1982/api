package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"my_project/controller"
	"net/http"
)

//var DB *sql.DB
var userCrt = &controller.UserCtrl{}

type UserDB struct {
	users controller.UserCtrl
}

func main() {
	DataSourceName := "user=fox password=123 dbname=fix sslmode=disable"
	Db, err := sql.Open("postgres", DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	us := &UserDB{
		users: controller.UserCtrl{DB: Db},
	}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", us.GetSingleUser).Methods("GET")
	router.HandleFunc("/users", us.Getusers).Methods("GET")
	router.HandleFunc("/user", us.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", us.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", us.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}
func (us *UserDB) Getusers(res http.ResponseWriter, req *http.Request) {
	us.Getusers(res, req)
}
func (us *UserDB) GetSingleUser(res http.ResponseWriter, req *http.Request) {
	us.GetSingleUser(res, req)
}
func (us *UserDB) CreateUser(res http.ResponseWriter, req *http.Request) {
	us.CreateUser(res, req)
}
func (us *UserDB) DeleteUser(res http.ResponseWriter, req *http.Request) {
	us.DeleteUser(res, req)
}
func (us *UserDB) UpdateUser(res http.ResponseWriter, req *http.Request) {
	us.UpdateUser(res, req)
}
