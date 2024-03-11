package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Online..."))
}