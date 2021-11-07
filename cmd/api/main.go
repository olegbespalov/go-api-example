package main

import (
	"log"
	"net/http"
	"time"

	"github.com/olegbespalov/go-api-example/pkg/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := config.AppPort
	if port == "" {
		log.Fatalln("port should be defined")
	}

	log.Printf("running an %s on http://localhost:%s/\n", config.AppName, port)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(config.MetricsAddr, nil))
	}()

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
