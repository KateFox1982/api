// Package controller, передача методов в Package model, согласно MVC
package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"strconv"

	"my_project/model"
	"net/http"
)

// UserCtrl структура используется для конструктора контроллер
type UserCtrl struct {
	users *model.UserModel
}

// NewUserCtrl конструктор контроллера, возращающий экземпляр структуры UserCtrl
// со свойством users контроллера модели с аргументом DB
func NewUserCtrl(DB *sql.DB) *UserCtrl {
	return &UserCtrl{
		users: model.NewUserModel(DB),
	}
}

// Getusers метод контроллера по получения всех значений из БД
func (usr *UserCtrl) Getusers(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//обращение к методу модели Getusers
	users, err := usr.users.Getusers()
	if err != nil {
		m := "Ошибка выполнеия функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&users)
}

// GetSingleUser метод контроллера по получению значения по id
func (usr *UserCtrl) GetSingleUser(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//изъятия из заголовка URL id string
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var id = params["id"]
	//конвертация string в int
	s, err := strconv.Atoi(id)
	if err != nil {
		m := "Ошибка перевода id из string в int "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//передача парметра id методу модели GetSingleUser
	p, err := usr.users.GetSingleUser(s)
	if err != nil {
		m := "Ошибка выполнения функции выбора по id: "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&p)

}

// CreateUser метод контроллера по созданию нового элемента в БД
func (usr *UserCtrl) CreateUser(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//объявление структуры User пакета model
	var user model.User
	//декорирование тела запроса в структуру
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		m := "Ошибка чтения информации для создания новой записи : "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//передача парметров запроса методу модели CreateUser
	m, err := usr.users.CreateUser(user.Name, user.Sale)
	if err != nil {
		m := "При выполнении функции создания возникла ошибка: "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&m)
}

// UpdateUser метод контроллера по изменению информации у конкретного id
func (usr *UserCtrl) UpdateUser(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//объявление структуры User пакета model
	var user model.User
	//декорирование тела запроса в структуру
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		m := "Ошибка маршаллинга данных для изменения"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//передача парметров запроса методу модели UpdateUser
	users, err := usr.users.UpdateUser(user.ID, user.Name, user.Sale)
	if err != nil {
		m := "При изменении что то пошло не так:"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&users)
}

// DeleteUser метод контроллера по удалению из БД по id
func (usr *UserCtrl) DeleteUser(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//изъятия из заголовка URL id string
	params := mux.Vars(req)
	id := params["id"]
	//конвертация string в int
	s, err := strconv.Atoi(id)
	if err != nil {
		m := "Неудачно выполнен перевод id bp string в int "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//передача парметра id методу модели GetSingleUser
	p, err := usr.users.DeleteUser(s)
	if err != nil {
		m := "Не удачно удалилось "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(p)
}
