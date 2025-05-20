package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
)

type Todo struct {
	UserID     int
	ID         int
	Title      string
	Compolated string
}

func performGetRequest() {
	gaurav := Todo{
		UserID:     23,
		ID:         1,
		Title:      "Buy milk",
		Compolated: "Buy milk",
	}
	//struct se json me convert kiya hu
	jsonData, err := json.Marshal(gaurav)

	if err != nil {
		fmt.Println(" Errors Marshalling", jsonData)
		return
	}
	//json data to string me convert kiya hu
	//jsonString := string(jsonData)
	//string data ko reader me convert kiya
	//jsonReader := string.NerReader(jsonString)

	//MyURL := "3fwfwfwf"

	//res, err := http.Post(MyURL, "application/json :", jsonString)
	// if err != nil {
	// 	fmt.Println("vsvs", res)

	// }
	// defer res.Body.Close()
	// data, _:=ioutil.ReadAll(res.Body)
	// fmt.Println("Respose ",data)

}

func main() {
	performGetRequest()
}
