package main

import (
	"breeders/adapters"
	"breeders/configuration"
	"breeders/streamer"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
	videoQueue  chan streamer.VideoProcessingJob
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	const numWorkers = 4

	videoQueue := make(chan streamer.VideoProcessingJob, numWorkers)
	defer close(videoQueue)

	app := application{
		templateMap: make(map[string]*template.Template),
		videoQueue:  videoQueue,
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(127.0.0.1:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// Get database
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// jsonBackend := &adapters.JSONBackend{}
	// jsonAdapter := &adapters.RemoteService{Remote: jsonBackend}

	xmlBackend := &adapters.XMLBackend{}
	xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	// app.App = configuration.New(db, jsonAdapter)
	app.App = configuration.New(db, xmlAdapter)

	wp := streamer.New(videoQueue, numWorkers)
	wp.Run()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
