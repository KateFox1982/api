package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"my_project/controller"
	"net/http"
)

type UserDB struct {
	users controller.UserCtrl
}

func main() {
	DataSourceName := "user=fox password=123 dbname=fix sslmode=disable"
	DB, err := sql.Open("postgres", DataSourceName)

	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	defer DB.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.Getusers(res, req)
	}).Methods("GET")
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.GetSingleUser(res, req)
	}).Methods("GET")
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.DeleteUser(res, req)
	}).Methods("DELETE")
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.UpdateUser(res, req)
	}).Methods("PUT")
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.CreateUser(res, req)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}
