package main

import (
	"log"
	"net/http"
	"github.com/guillaumeparis2000/rest-api/models"
)

func main(){
	router := NewRouter()

	models.InitDb()

	log.Fatal(http.ListenAndServe(":8080", router))
}
