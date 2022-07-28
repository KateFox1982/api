//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//
//	"net/http"
//)
//
////
////type User struct {
////	ID   int    `json:"id"`
////	Name string `json:"name"`
////	Sale int    `json:"sale"`
////}
//
//func MakeRequest() {
//
//	resp, err := http.Get("https://127.0.0.1:4000/users")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//
//	io.Copy(os.Stdout, resp.Body)
//
//}

//	conn, err := net.Dial("tcp", "127.0.0.1:4545")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer conn.Close()
//	resp, err := http.Get("https://127.0.0.1:4545/users", "Get")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	for true {
//
//		bs := make([]byte, 1014)
//		n, err := resp.Body.Read(bs)
//		fmt.Println(string(bs[:n]))
//		if n == 0 || err != nil {
//			break
//		}
//
//		var users User
//		err := json.Unmarshal(bs, &users)
//		if err != nil {
//			log.Println(err)
//		}
//		map:=make ([] int)
//map:=append(map, users.ID)
//i:=range (map)
//id :=map[i-1]
//		resp, err := http.Post("https://127.0.0.1:4545/user/{id}", "DELETE")
//		if err != nil {
//		fmt.Println(err)
//		return
//		}
//defer resp.Body.Close()
//for true {
//
//bs := make([]byte, 1014)
//n, err := resp.Body.Read(bs)
//fmt.Println(string(bs[:n]))
//if n == 0 || err != nil {
//break
//}
//id=map[0]
//resp, err := http.Post("https://127.0.0.1:4545/user/{id}", "PUT")
//if err != nil {
//fmt.Println(err)
//return
//}
//defer resp.Body.Close()
//for true {
//w.Header().Set("Content-Type", "application/json")
//var user User
//users=append(users, User{Name: "Sofa", Sale: 543})
//_ = json.NewDecoder(r.Body).Decode(&user)
//user.ID = strconv.Itoa(rand.Intn(1000000))
//users = append(users, user)
//json.NewEncoder(w).Encode(user)
//
//bs := make([]byte, 1014)
//n, err := resp.Body.Read(bs)
//fmt.Println(string(bs[:n]))
//if n == 0 || err != nil {
//break
//}
//resp, err := http.Post("https://127.0.0.1:4545/user", "POST")
//if err != nil {
//fmt.Println(err)
//return
//}
//defer resp.Body.Close()
//for true {
//
//bs := make([]byte, 1014)
//n, err := resp.Body.Read(bs)
//fmt.Println(string(bs[:n]))
//if n == 0 || err != nil {
//break
//}
//resp, err := http.Post("https://127.0.0.1:4545/users", "Get")
//if err != nil {
//fmt.Println(err)
//return
//}
//defer resp.Body.Close()
//for true {
//
//bs := make([]byte, 1014)
//n, err := resp.Body.Read(bs)
//fmt.Println(string(bs[:n]))
//if n == 0 || err != nil {
//break
//}
//
//	}
//}
