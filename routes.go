package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-web-native/controllers/homecontroller"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addApproutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/", homecontroller.Welcome)
	route.HandleFunc("/user", homecontroller.Insert).Methods("Post")

	fmt.Println("Routes are Loded.")
}
