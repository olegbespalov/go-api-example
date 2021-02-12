package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/go-api-example/pkg/config"
)

func main() {
	port := config.AppPort

	log.Printf("running an API on http://localhost:%s/\n", port)

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicln(err.Error())
	}
}
