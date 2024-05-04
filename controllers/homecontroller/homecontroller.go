package homecontroller

import (
	// "go-web-native/models/personModel"

	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

type User struct {
	firstName string
	lastName  string
}

func Insert(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	url := r.URL.String()
	queryParams := r.URL.Query()
	fmt.Println(method, url, queryParams)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := r.Form.Get("firstName")
	email := r.Form.Get("lastName")
	fmt.Println(name, email)

	var rlt = make(map[string]string)
	rlt["brand"] = "Ford"
	rlt["model"] = "Mustang"
	rlt["year"] = "1964"
	fmt.Println(rlt)
	jsonData, err := json.Marshal(rlt)
	delete(rlt, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
