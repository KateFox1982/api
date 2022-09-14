package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"strconv"

	"my_project/model"
	"net/http"
)

type UserCtrl struct {
	users *model.UserModel
}

//конструктор контроллера
func NewUserCtrl(DB *sql.DB) *UserCtrl {
	return &UserCtrl{
		users: &model.UserModel{
			DB: DB},
	}
}

//метод контроллера по получения всех значений из БД
func (usr *UserCtrl) Getusers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	users, err := usr.users.Getusers()
	if err != nil {
		log.Printf("Ошибка выполнеия функции получения информации о всех пользователях: %s", err)
		return
	}
	json.NewEncoder(res).Encode(&users)
}

//метод контроллера по получению значения по id
func (usr *UserCtrl) GetSingleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var id = params["id"]
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Ошибка перевода id из string в int %s", err)
		return
	}

	p, err := usr.users.GetSingleUser(s)
	if err != nil {
		log.Printf("Ошибка выполнения функции выбора по id: %s", err)
		return
	}

	json.NewEncoder(res).Encode(&p)

}

//метод контроллера по созданию нового элемента в БД
func (usr *UserCtrl) CreateUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {

		log.Printf("Ошибка чтения информации для сздания новой записи : %s", err)
		return
	}
	m, err := usr.users.CreateUser(user.Name, user.Sale)
	if err != nil {
		log.Printf("При выполнении функции создания возникла ошибка: %s", err)
		return
	}
	json.NewEncoder(res).Encode(&m)
}

//метод контроллера по изменению информации у конкретного id
func (usr *UserCtrl) UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Fatal("Ошибка маршаллинга данных для изменения")
		return
	}
	fmt.Println(user)
	fmt.Println(user.ID, user.Name, user.Sale)
	users, err := usr.users.UpdateUser(user.ID, user.Name, user.Sale)
	if err != nil {
		log.Printf("При изменении что то пошло не так: %s", err)
		return
	}

	json.NewEncoder(res).Encode(&users)
}

//метод контроллера по удалению из БД по id
func (usr *UserCtrl) DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Неудачно выполнен перевод id bp string в int %s", err)
	}
	p, err := usr.users.DeleteUser(s)
	if err != nil {
		log.Printf("Все удачно удалилось %s", id)
		return
	}
	json.NewEncoder(res).Encode(p)
}
