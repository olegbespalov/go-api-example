package main

import (
	"log"
	"net/http"
	"time"

	"github.com/olegbespalov/go-api-example/pkg/config"
)

func main() {
	port := config.AppPort
	if port == "" {
		log.Fatalln("port should be defined")
	}

	log.Printf("running an %s on http://localhost:%s/\n", config.AppName, port)

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
