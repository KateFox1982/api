package model

import (
	"database/sql"
	"fmt"
)

// Document структура используется инициализации данные в структуры
type Document struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

// DocumentModel используется для конструктора модели
type DocumentModel struct {
	DB *sql.DB
}

// NewUserModel конструктор модели возвращающий указатель на структуру UserModel
func NewDocumentModel(DB *sql.DB) *DocumentModel {
	return &DocumentModel{
		DB: DB,
	}
}

// GetDocuments метод модели по получению всех пользователей из БД возвращает массив структур Document и ошибку
func (m *DocumentModel) GetDocuments() ([]Document, error) {
	//rows запрос возврата срок выборки из таблицы значений
	var rows, err = m.DB.Query("SELECT id, title FROM documentations.document")
	if err != nil {
		fmt.Println("Ошибка в выбора таблицы ", err)
		return nil, err
	}
	defer rows.Close()
	//document инициализация массива структур Document
	document := []Document{}
	//получение данных из всей таблицы
	for rows.Next() {
		p := Document{}
		err := rows.Scan(&p.Id, &p.Title)
		if err != nil {
			fmt.Println("Ошибка сканирования результата селекта ", err)
			return nil, err
		}
		//добавление новых данных в массив структур
		document = append(document, p)
	}
	//возврат массива структур и ошибки
	return document, err
}
