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

// структура UserCtrl
type UserCtrl struct {
	users *model.UserModel
}

//конструктор контроллера, возращающий экземпляр структуры UserCtrl
// со свойством users контроллера модели с аргументом DB
func NewUserCtrl(DB *sql.DB) *UserCtrl {
	return &UserCtrl{
		users: model.NewUserModel(DB),
		//users: &model.UserModel{
		//	DB: DB},
	}
}

//метод контроллера по получения всех значений из БД
func (usr *UserCtrl) Getusers(res http.ResponseWriter, req *http.Request) {
	//прием заголовка URL, парсинг в json
	res.Header().Set("Content-Type", "application/json")
	//обращение к методу модели Getusers
	users, err := usr.users.Getusers()
	if err != nil {
		log.Printf("Ошибка выполнеия функции получения информации о всех пользователях: %s", err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&users)
}

//метод контроллера по получению значения по id
func (usr *UserCtrl) GetSingleUser(res http.ResponseWriter, req *http.Request) {
	//прием заголовка URL, парсинг в json
	res.Header().Set("Content-Type", "application/json")
	//изъятия из заголовка URL id string
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var id = params["id"]
	//конвертация string в int
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Ошибка перевода id из string в int %s", err)
		return
	}
	//передача парметра id методу модели GetSingleUser
	p, err := usr.users.GetSingleUser(s)
	if err != nil {
		log.Printf("Ошибка выполнения функции выбора по id: %s", err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&p)

}

//метод контроллера по созданию нового элемента в БД
func (usr *UserCtrl) CreateUser(res http.ResponseWriter, req *http.Request) {
	//прием заголовка URL, парсинг в json
	res.Header().Set("Content-Type", "application/json")
	//объявление структуры User пакета model
	var user model.User
	//декорирование тела запроса в структуру
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {

		log.Printf("Ошибка чтения информации для сздания новой записи : %s", err)
		return
	}
	//передача парметров запроса методу модели CreateUser
	m, err := usr.users.CreateUser(user.Name, user.Sale)
	if err != nil {
		log.Printf("При выполнении функции создания возникла ошибка: %s", err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&m)
}

//метод контроллера по изменению информации у конкретного id
func (usr *UserCtrl) UpdateUser(res http.ResponseWriter, req *http.Request) {
	//прием заголовка URL, парсинг в json
	res.Header().Set("Content-Type", "application/json")
	//объявление структуры User пакета model
	var user model.User
	//декорирование тела запроса в структуру
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Fatal("Ошибка маршаллинга данных для изменения")
		return
	}
	//передача парметров запроса методу модели UpdateUser
	users, err := usr.users.UpdateUser(user.ID, user.Name, user.Sale)
	if err != nil {
		log.Printf("При изменении что то пошло не так: %s", err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&users)
}

//метод контроллера по удалению из БД по id
func (usr *UserCtrl) DeleteUser(res http.ResponseWriter, req *http.Request) {
	//прием заголовка URL, парсинг в json
	res.Header().Set("Content-Type", "application/json")
	//изъятия из заголовка URL id string
	params := mux.Vars(req)
	id := params["id"]
	//конвертация string в int
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Неудачно выполнен перевод id bp string в int %s", err)
	}
	//передача парметра id методу модели GetSingleUser
	p, err := usr.users.DeleteUser(s)
	if err != nil {
		log.Printf("Все удачно удалилось %s", id)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(p)
}
