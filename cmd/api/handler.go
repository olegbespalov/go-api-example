package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/go-api-example/pkg/response"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request")

	response.Message(w, http.StatusOK, "Hey Here")
}
