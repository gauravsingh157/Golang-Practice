package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	users = make(map[string]User)
)

func adduser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err :=ErrorRes{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		json.NewEncoder(w).Encode(err)

		return
	}

	users[user.Name] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func addGetuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err :=json.NewEncoder(w).Encode(users)
	if err!= nil{
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusBadRequest)
		err :=ErrorRes{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}

		json.NewEncoder(w).Encode(err)
	
	}
}




func main() {
	fmt.Println("Server Started ........")
	http.HandleFunc("/GAURAV", adduser)
	http.HandleFunc("/SINGH",addGetuser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
