package main

import (
	"awesomeProject/handler"
	"awesomeProject/service"
	"awesomeProject/store"
	"log"
	"net/http"
)

func main() { 
	str := store.New()
	svc := service.New(str)
	hld := handler.New(svc)

	http.HandleFunc("/account", hld.Router)
	log.Println("listen on 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
