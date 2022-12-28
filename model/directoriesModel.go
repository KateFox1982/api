package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//постоянная с адресом из какой директории надо получить информацию о вложенности
const path = "/home/kate/Music"

//Directory структура имеющая рекурсионный вид
type Directory struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Directories []Directory
}

//DirectoriesModel используется для конструктора модели
type DirectoriesModel struct {
	HTTPServer http.Server
}

// NewDirectoriesModel конструктор модели
func NewDirectoriesModel() *DirectoriesModel {
	return &DirectoriesModel{HTTPServer: http.Server{}}
}

// GetDirectories метод модели получающий слайс Директорий находящихся по адресу path
func (dir *DirectoriesModel) GetDirectories(path string) (*[]Directory, error) {
	//создание экземпляра структуры Directory
	directories := []Directory{}
	var fistDirectoryTitle string
	//разбиение path на слова по разделителю "/"
	str := strings.Split(path, "/")
	num := 1
	//длина строки path
	lastDir := len(str)
	//цикл для определения "первоначальной" диретории
	for i := 0; i < lastDir; i++ {
		fistDirectoryTitle = ``
		fistDirectory := str[lastDir-1]
		fistDirectoryTitle = fistDirectoryTitle + fistDirectory
	}
	//инициализация данных, полученых выше в структуру Directory
	directory := Directory{Id: num, Title: fistDirectoryTitle}
	//обращение к рекурсионной функции readDir
	err := readDir(path, &directory.Directories, num)
	if err != nil {
		fmt.Println("Ошибка извлечения директорий", err)
		return nil, err
	}
	//добавление структуры в слайс структур
	directories = append(directories, directory)
	return &directories, nil
}

//readDir рекурсионная функция необходимая длянахождения всех вложенных Диреторий
func readDir(path string, directories *[]Directory, num int) error {
	//чиение диреторий находящихся в path, в случае если бы необходимо
	//прочитать вместе с файлами нужно использовать Reddirnames
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	directory := Directory{}
	//цикл для изменения адреса path, определения является ли объект директорией или файлом,
	//а так е рексурсионным вызовом самой себя, для получения вложенных объектов
	for _, file := range files {
		path := path + `/` + file.Name()
		fmt.Println("path ", path)
		if file.IsDir() == true {
			num = num + 1
			fmt.Println("file.Name()", file.Name())
			directory = Directory{Id: num, Title: file.Name()}
			fmt.Println("directory", directory)
			err = readDir(path, &directory.Directories, num)
			fmt.Println("directory.Directories", &directory.Directories)
			//directories = append(directories, directory)
			fmt.Println("directories", directories)
			*directories = append(*directories, directory)
		}
	}
	fmt.Println("directories", directories)
	return err
}
