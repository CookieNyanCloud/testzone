package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	router := mux.NewRouter()
	const post string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_,_ =fmt.Fprintln(resp,"Up and running...")
	})
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	log.Println("server listening on port ", post)
	log.Fatalln(	http.ListenAndServe(post,router))
}

//todo:not over