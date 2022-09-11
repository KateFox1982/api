package model

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sale int    `json:"sale"`
}
type UserModel struct {
	DB *sql.DB
}

//func NewModel(id int, name string, sale int) *Model {
//
//	return &Model{
//		id,
//		name,
//		sale,
//	}

func (m *UserModel) Getusers() ([]User, error) {
	var rows, err = m.DB.Query("SELECT * FROM Misha2")
	if err != nil {
		log.Fatal(err)
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
	return users, err
}
func (m *UserModel) GetSingleUser(id int) (User, error) {
	var p User

	row1 := m.DB.QueryRow("SELECT id, name, sale FROM Misha2 where id=$1", id)

	err := row1.Scan(&p.ID, &p.Name, &p.Sale)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("Нет такого id=%d", id)
		return p, err
		//	//else if err != nil {
		//	//	fmt.Println("Unexpected error: ", err.Error())
	} else {
		return p, nil
	}
	return p, err
}
func (m *UserModel) CreateUser(name string, sale int) (User, error) {

	var user *User

	//var lastInsertID int

	err := m.DB.QueryRow("INSERT INTO Misha2 (name, sale) VALUES($1,$2) returning id", user.Name, user.Sale).Scan(&user.ID)
	if err != nil {

		fmt.Println(err)
		return User{}, err
	}
	return User{}, nil
}
func (m *UserModel) UpdateUser(id int, name string, sale int) (User, error) {
	//
	var user *User
	_, err := m.DB.Exec("update Misha2 set name = $1, sale= $2 where id = $3", user.Name, user.Sale, user.ID)
	if err != nil {
		err = fmt.Errorf("Нет такого id=%d", id)
		return User{}, err
	}

	return User{}, err
}
func (m *UserModel) DeleteUser(id int) (User, error) {

	var s *string

	row1 := m.DB.QueryRow("SELECT FROM Misha2 where id=$1", id)
	//p := new(ID)

	err := row1.Scan(&s)

	if err == sql.ErrNoRows {
		err = fmt.Errorf("Нет такого id=%d", id)
		return User{}, err
	} else {
		var k string
		row := m.DB.QueryRow("DELETE  FROM Misha2 where id=$1", id)

		err = row.Scan(&k)
		if err == sql.ErrNoRows {
			fmt.Println()
		}
	}
	return User{}, err
}
