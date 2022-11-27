package model

import (
	"database/sql"
	fmt "fmt"
)

// Document структура используется инициализации данные в структуры
type Document struct {
	Id      int64    `json:"id" :"id"`
	Title   string   `json:"title" :"title"`
	Modules []Module `json:"modules"`
}

// DocumentModel используется для конструктора модели
type DocumentModel struct {
	dataBase    *sql.DB
	moduleModel *ModuleModel
	errorModel  *ErrorModel
}

// NewUserModel конструктор модели возвращающий указатель на структуру DocumentModel
func NewDocumentModel(DB *sql.DB) *DocumentModel {
	return &DocumentModel{
		dataBase:    DB,
		moduleModel: NewModuleModel(DB),
		errorModel:  NewErrorModel(DB),
	}
}

// GetDocuments метод модели по получению всех пользователей из БД возвращает массив структур Document и ошибку
func (m *DocumentModel) GetDocuments() ([]Document, error) {
	//rows запрос возврата срок выборки из таблицы значений
	var rows, err = m.dataBase.Query("SELECT id, title FROM documentations.document")
	if err != nil {
		err := fmt.Errorf("Ошибка в выбора таблицы %s ", err)
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
			err := fmt.Errorf("ошибка сканирования результата селекта %s", err)
			return nil, err
		}
		//добавление новых данных в массив структур
		document = append(document, p)
	}
	//возврат массива структур и ошибки
	return document, err
}
func (m *DocumentModel) GetDocumentsFull() ([]Document, error) {
	var doc = []Document{}
	doc, err := m.GetDocuments()
	if err != nil {
		err := fmt.Errorf("ошибка сканирования результата селекта %s", err)
		return []Document{}, err
	}
	for i := range doc {
		docId := doc[i].Id
		doc[i].Modules, err = m.moduleModel.GetModuleById(docId)
		if err != nil {
			err := fmt.Errorf("ошибка сканирования результата селекта %s", err)
			return []Document{}, err
		}
		for k := range doc[i].Modules {

			moduleId := doc[i].Modules[k].Id
			//var errorModel []Error
			doc[i].Modules[k].Errors, err = m.errorModel.GetErrorById(moduleId)
			if err != nil {
				err := fmt.Errorf("ошибка сканирования результата селекта %s", err)
				return []Document{}, err
			}

		}
	}

	return doc, err
}
