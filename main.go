package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("this is simple server"))
	})
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error")
	})
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthcheck"))
	})
	fmt.Println("http server is running on port 8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}