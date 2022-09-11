package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func OpenConnection() *gorm.DB {
	dataSourceName := "user=fox password=123 dbname=fix sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
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
	db.AutoMigrate(&Track{})

	return db
}
