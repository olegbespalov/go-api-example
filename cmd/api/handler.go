package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/go-api-example/pkg/response"
)

type simple struct{}

func (s simple) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	log.Println("incoming request")

	response.Message(w, http.StatusOK, "Hey Here")
}
