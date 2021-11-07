package main

import (
	"log"
	"net/http"
	"time"

	"github.com/olegbespalov/go-api-example/pkg/config"
)

func main() {
	port := config.AppPort

	log.Printf("running an API on http://localhost:%s/\n", port)

	srv := http.Server{
		Addr:         ":" + port,
		Handler:      &simple{},
		ReadTimeout:  2500 * time.Millisecond,
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err.Error())
	}
}
