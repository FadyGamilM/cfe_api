package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/cfe_api/db"
)

// global variable
var server_port = os.Getenv("PORT")

// configurations for the server
type AppConfig struct {
	// keep it uppercase so i can use it into other packages
	PORT string
}

// represents the entire server
type Application struct {
	Config AppConfig
}

func (app *Application) Serve() error {
	// instantiate a http server instance
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", server_port),
	}

	// return a running server instance or an error
	return server.ListenAndServe()
}

func main() {
	server_configs := AppConfig{PORT: server_port}
	dsn := os.Getenv("DSN")

	dbConnPool, dbConnErr := db.ConnectToPostgresInstance(dsn)
	if dbConnErr != nil {
		log.Fatalf("cannot connect to the db instance \n ERROR ➜ %v \n", dbConnErr)
	}
	// if connection goes well, we need to defer the db closing
	defer dbConnPool.DB.Close()

	// create instance from the Application and start the server
	app := &Application{Config: server_configs}
	serverErr := app.Serve()
	if serverErr != nil {
		log.Fatalf("cannot start the server \n ERROR ➜ %v \n", serverErr)
	}
}
