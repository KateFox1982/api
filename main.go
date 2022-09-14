package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"my_project/controller"
	"net/http"
)

func main() {
	//парметры БЛ: имя пользователя, пароль, имя БД, использование SSL
	DataSourceName := "user=fox password=123 dbname=fix sslmode=disable"
	//соединение с БД postgres
	DB, err := sql.Open("postgres", DataSourceName)
	//ошибка соединения
	if err != nil {
		log.Printf("Got error in mysql connector: %s", err)
		return
	}
	//отсрочка закрытия БД
	defer DB.Close()
	//запуск роутера
	router := mux.NewRouter()
	//регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.Getusers(res, req)
	}).Methods("GET")
	//регистрация второго маршрута, с URL оканчивающимся на "/user и параметром id, который пользователь указывает в URL,
	//и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера GetSingleUser
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.GetSingleUser(res, req)
	}).Methods("GET")
	//регистрация третьего маршрута, с URL оканчивающимся на "/user и параметром id, который пользователь указывает в URL,
	//и методом DELETE, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера DeleteUser
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.DeleteUser(res, req)
	}).Methods("DELETE")
	//регистрация третьего маршрута, с URL оканчивающимся на "/user ,
	//и методом PUT, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера UpdateUser
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.UpdateUser(res, req)
	}).Methods("PUT")
	//регистрация третьего маршрута, с URL оканчивающимся на "/user ,
	//и методом POST, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера CreateUser
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserCtrl(DB)
		userCtrl.CreateUser(res, req)
	}).Methods("POST")
	//прослушивание браузера
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}
