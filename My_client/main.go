package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"

	//"io"
	"io/ioutil"
	//"log"
	"net/http"
)
type User struct {
ID   int    `xml:"id_xml",json:"id"`
Name string `xml:"name_xml",json:"name"`
Sale int    `xml:"sale_xml",json:"sale_xml"`
}


func main() {

	router := mux.NewRouter()
	router.HandleFunc("/do", HandleHttp).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
	}
func HandleHttp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Сервер запустился")
	//MakeRequestGet()
	user, err := MakeRequestGet()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Переданная структура", user)
	idMin := Min(user)
	if err != nil {
		fmt.Println(err)
	}
	idMax := Max(user)
	fmt.Println("Максимальное значение", idMax)
	fmt.Println("Минимальное значение", idMin)

	MakeRequestUpdate(idMin)
	MakeRequestDelete(idMax)
	MakeRequestCreate()
	t, err:=MakeRequestGet()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Срез структуры перед XML", t)
	res.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(res).Encode(&t)
}
func MakeRequestUpdate (idMin int) (User, error) {
client := http.Client{}
var user User
user=User{
ID: idMin,
Name: "Vova",
Sale: 654,
}
userBytes, err:= json.Marshal(user)
if err != nil {
fmt.Println(err)
}
byteRead:=bytes.NewReader(userBytes)
req, err:= http.NewRequest("PUT", "http://localhost:4000/user", byteRead)

if err != nil {
fmt.Println(err)
}
resp, err := client.Do(req)
if err != nil {
fmt.Println(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body) // response body is []byte
if err !=nil {
fmt.Println(err)
}
fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, fmt.Errorf("can't parse body as JSON: %w", err)
	}
	return user, err
}
func MakeRequestDelete (idMax int) (User, error){
client := http.Client{}

id:= strconv.Itoa(idMax)
fmt.Println("Максимально id",id)
id=url.PathEscape(id)


u, err := url.Parse("http://127.0.0.1:4000/user/")
if err != nil {
log.Fatal(err)
}

rel, err := u.Parse(id)
if err != nil {
log.Fatal(err)
}
fmt.Println(rel)

udlStr:=rel.String()

req, err:= http.NewRequest("DELETE", udlStr, nil)

if err != nil {
fmt.Println(err)

}
resp, err := client.Do(req)
if err != nil {
fmt.Println(err)

}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body) // response body is []byte
if err !=nil {
fmt.Println(err)

}
fmt.Println(string(body))


//fmt.Println(PrettyPrint(result))

// Loop through the data node for the FirstName
io.Copy(os.Stdout, resp.Body)
	return User{}, err
}
func MakeRequestGet() (p []User, err error) {
client := http.Client{}
req, err := http.NewRequest("GET", "http://127.0.0.1:4000/users", nil)
if err != nil {
fmt.Println("Проблема с адресом", err)
}
resp, err := client.Do(req)
if err != nil {
fmt.Println("Проблема с ответом", err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body) // response body is []byte
if err != nil {
fmt.Println(err)
}
p=[]User{}

fmt.Println("Печать из функции", string(body))
err = json.Unmarshal(body, &p)
if err != nil {
fmt.Println("Can not unmarshal JSON", err)
}
fmt.Println("Структура", p)
return
}
func Min (p []User) int {

var m []int

for _, rec := range p {
m= append(m, rec.ID)
}
idMin:=m[0]
for _, value := range m {
if value < idMin {
idMin = value
}
}
return idMin
}
func Max (p []User)  int {




var m []int

for _, rec := range p {
m= append(m, rec.ID)
}

idMax:=m[0]

for _, value := range m {
if value > idMax {
idMax = value
}
}
fmt.Println(idMax)
return idMax
}
func MakeRequestCreate () (User, error) {
client := http.Client{}
var user User
user=User{
Name: "RED",
Sale: 895,
}
userBytes, err:= json.Marshal(user)
if err != nil {
fmt.Println(err)

}
byteRead:=bytes.NewReader(userBytes)
req, err:= http.NewRequest("POST", "http://127.0.0.1:4000/user", byteRead)
if err != nil {
fmt.Println(err)
}
resp, err := client.Do(req)
if err != nil {
fmt.Println(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body) // response body is []byte
if err !=nil {
fmt.Println(err)
}
fmt.Println(string(body))
	return User{}, err
}
