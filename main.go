package main

import (
	"log"
	"net/http"
	"github.com/guillaumeparis2000/rest-api/models"
	"github.com/guillaumeparis2000/rest-api/router"
)

func main(){
	router := router.NewRouter()

	models.InitDb()

	log.Fatal(http.ListenAndServe(":8080", router))
}
