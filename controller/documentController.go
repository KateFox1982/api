package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"my_project/model"
	"net/http"
)

// DocumentController структура используется для конструктора контроллер
type DocumentController struct {
	model *model.DocumentModel
}

// NewDocumentModel конструктор контроллера, возращающий экземпляр структуры DocumentController
// со свойством users контроллера модели с аргументом DB
func NewDocumentModel(DB *sql.DB) *DocumentController {
	return &DocumentController{
		model: model.NewDocumentModel(DB),
	}
}

// GetDocuments метод контроллера по получения всех значений из БД
func (dc *DocumentController) GetDocuments(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	//обращение к методу модели Getusers
	documents, err := dc.model.GetDocuments()
	if err != nil {
		m := "Ошибка выполнеия функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&documents)
}
