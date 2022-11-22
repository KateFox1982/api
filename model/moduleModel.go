package model

import (
	"database/sql"
	"fmt"
)

// Module структура используется инициализации данные в структуры
type Module struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

// ModuleModel используется для конструктора модели
type ModuleModel struct {
	DB *sql.DB
}

// NewModuleModel конструктор модели возвращающий указатель на структуру UserModel
func NewModuleModel(DB *sql.DB) *ModuleModel {
	return &ModuleModel{
		DB: DB,
	}
}

// GetModuleById метод модели по получению всех пользователей из БД возвращает массив структур Module по id документа и ошибку
func (m *ModuleModel) GetModuleById(documentId int64) ([]Module, error) {
	//QueryRow запрос возврата сроки выборки из таблицы значений значений по id
	var rows, err = m.DB.Query("SELECT id, title FROM documentations.module where fk_document=$1", documentId)
	if err != nil {
		fmt.Println("Ошибка в выбора таблицы ", err)
		return nil, err
	}
	defer rows.Close()
	module := []Module{}
	//получение данных из всей таблицы
	for rows.Next() {
		p := Module{}
		err := rows.Scan(&p.Id, &p.Title)
		if err != nil {
			fmt.Println("Ошибка сканирования результата селекта ", err)
			return nil, err
		}
		//добавление новых данных в массив структур
		module = append(module, p)
	}
	return module, err
}
