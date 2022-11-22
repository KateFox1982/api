//Package main входная точка АПИ, запуск роутеров
package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"my_project/controller"
	"net/http"
)

func main() {
	//парметры БД: имя пользователя, пароль, имя БД, использование SSL
	DataSourceName := "user=fox password=123 dbname=fix sslmode=disable "
	//соединение с БД postgres
	DB, err := sql.Open("postgres", DataSourceName)
	//err ошибка соединения
	if err != nil {
		log.Printf("Получить ошибку о postgres присоединении: %s", err)
		return
	}
	//defer Close отсрочка закрытия БД
	defer DB.Close()
	//запуск роутера
	router := mux.NewRouter()
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		userCtrl := controller.NewUserController(DB)
		userCtrl.Getusers(res, req)
	}).Methods("GET")
	//router.HandleFunc регистрация второго маршрута, с URL оканчивающимся на "/user и параметром id, который пользователь указывает в URL,
	//и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера GetSingleUser
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserController(DB)
		userCtrl.GetSingleUser(res, req)
	}).Methods("GET")
	//router.HandleFunc регистрация третьего маршрута, с URL оканчивающимся на "/user и параметром id, который пользователь указывает в URL,
	//и методом DELETE, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера DeleteUser
	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserController(DB)
		userCtrl.DeleteUser(res, req)
	}).Methods("DELETE")
	//router.HandleFunc регистрация третьего маршрута, с URL оканчивающимся на "/user,
	//и методом PUT, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера UpdateUser
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserController(DB)
		userCtrl.UpdateUser(res, req)
	}).Methods("PUT")
	//router.HandleFunc регистрация третьего маршрута, с URL оканчивающимся на "/user ,
	//и методом POST, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера CreateUser
	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewUserController(DB)
		userCtrl.CreateUser(res, req)
	}).Methods("POST")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/documents" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/documents", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		userCtrl := controller.NewDocumentModel(DB)
		userCtrl.GetDocuments(res, req)
	}).Methods("GET")
	//router.HandleFunc регистрация второго маршрута, с URL оканчивающимся на "/module и параметром id, который пользователь указывает в URL,
	//и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера GetModuleById
	router.HandleFunc("/module/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewModuleModel(DB)
		userCtrl.GetModuleById(res, req)
	}).Methods("GET")
	//router.HandleFunc регистрация второго маршрута, с URL оканчивающимся на "/error и параметром id, который пользователь указывает в URL,
	//и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера GetErrorById
	router.HandleFunc("/error/{id}", func(res http.ResponseWriter, req *http.Request) {
		userCtrl := controller.NewErrorModel(DB)
		userCtrl.GetErrorById(res, req)
	}).Methods("GET")
	//ListenAndServe запуск http -сервером с адресом 127.0.0.1:4000, и обработчиком router
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}
