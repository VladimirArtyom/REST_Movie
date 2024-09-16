package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env string
}

type application struct {
	config config  
	logger *log.Logger 
} 

func main() {

	// Define the config object
	// Define flags for the config values
	// Parse the flags
	// Define Log object
	// Integrate it to application

	var cfg config
	
	flag.IntVar(&cfg.port, "port", 8080, "API Default server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (devevelopment|staging|production)")
	flag.Parse()

	var logger *log.Logger = log.New(os.Stdout, "",  log.Ldate | log.Ltime)

	var app *application = &application{
		config: cfg,
		logger: logger,
	}

	// Make a multiplexer, basically request handler using mux basically

	var server *http.Server = &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}


	logger.Printf("Starting server on %s in %s mode", server.Addr, cfg.env)
	server.ListenAndServe()

	fmt.Println("Hello world")
}
