package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sale int    `json:"sale"`
}

var users []User

func OpenConnection() *sql.DB {
	dataSourceName := "user=fox password=123 dbname=fix sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//_, err = db.Exec("DROP TABLE  IF EXISTS Misha2")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//insertTab := `create table IF NOT EXISTS  Misha2 (id SERIAL, Name TEXT NOT NULL, Sale INT, CONSTRAINT Mish2_pkey PRIMARY KEY (id)) `
	//_, err = db.Exec(insertTab)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = db.Exec("ALTER SEQUENCE Misha2_id_seq RESTART WITH 1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = db.Exec("TRUNCATE  Misha2")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = db.Exec("insert into  Misha2 (id, Name, Sale) values(1, 'Kate', $1)", 897)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec("insert into  Misha2 (id, Name, Sale) values(2, 'Kate', 89)")
	//if err != nil {
	//	panic(err)
	//	//fmt.Println("Ошибка 1", err)
	//	//return
	//}
	//
	//_, err = db.Exec("insert into  Misha2 (id, Name, Sale) values(3, 'Jhon', $1)", 768)
	//if err != nil {
	//	//fmt.Println("Ошибка 2", err)
	//	//return
	//}
	//_, err = db.Exec("insert into  Misha2 (id, Name, Sale) values(4, 'Lera', $1)", 546)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec("insert into  Misha2 (id, Name, Sale) values(5, 'Lada', $1)", 678)
	//if err != nil {
	//	panic(err)
	//}
	return db
}

func getusers(res http.ResponseWriter, req *http.Request) {
	// we have to set the header "Content-Type: application/json"
	// because we are sending JSON data with a request through postman
	res.Header().Set("Content-Type", "application/json")
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM Misha2")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		p := User{}

		err := rows.Scan(&p.ID, &p.Name, &p.Sale)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {
		fmt.Println(p.ID, p.Name, p.Sale)
	}
	// we are taking variable 'courses' in which we've appended dummy data and returning that as a response
	json.NewEncoder(res).Encode(&users)
}

func getSingleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	db := OpenConnection()
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url
	id := params["id"]
	var p User

	row1 := db.QueryRow("SELECT id, name, sale FROM Misha2 where id=$1", id)

	err := row1.Scan(&p.ID, &p.Name, &p.Sale)

	if err == sql.ErrNoRows {
		//fmt.Println(err.Error())
		fmt.Fprintf(res, "Нет такого id=%s", id)
		//	//else if err != nil {
		//	//	fmt.Println("Unexpected error: ", err.Error())
	} else {

		json.NewEncoder(res).Encode(&p)

		//json.NewEncoder(res).Encode("No course found")
	}

}
func createUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//
	var user *User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	//Name := req.FormValue("Name")
	//fmt.Printf(" N=%s ", Name)
	//
	//Sale := req.FormValue("Sale")
	//fmt.Printf(" S=%s ", Sale)
	//if Name == "" || Sale == "" {
	//	fmt.Println(" Пусто")
	//
	//}
	//fmt.Println("Вставляем  %s" + Name + " and sale: " + Sale)

	db := OpenConnection()
	//var lastInsertID int

	err = db.QueryRow("INSERT INTO Misha2 (name, sale) VALUES($1,$2) returning id", user.Name, user.Sale).Scan(&user.ID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(&user)
}
func updateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//
	var user *User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	db := OpenConnection()
	_, err = db.Exec("update Misha2 set name = $1, sale= $2 where id = $3", user.Name, user.Sale, user.ID)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(res).Encode(&user)

}
func deleteUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	db := OpenConnection()
	id := params["id"]

	var s *string

	row1 := db.QueryRow("SELECT FROM Misha2 where id=$1", id)
	//p := new(ID)

	err := row1.Scan(&s)
	//if err != nil {
	//	fmt.Fprintf(res, "Нет такого  id=%s ", id)
	//	return
	//} else {
	//	fmt.Fprintf(res, "Хрен его знает что=%s ", id)
	//
	//	return
	//}

	if err == sql.ErrNoRows {
		fmt.Fprintf(res, "Нет такого id=%s ", id)

	} else {
		var k string
		row := db.QueryRow("DELETE  FROM Misha2 where id=$1", id)

		err = row.Scan(&k)
		if err == sql.ErrNoRows {
			fmt.Println()
		}

		json.NewEncoder(res).Encode(k)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", getSingleUser).Methods("GET")
	router.HandleFunc("/users", getusers).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}
