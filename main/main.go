package main

import (
	cc "github.com/rashmi43/go-messenger/controller"
	"github.com/rashmi43/go-messenger/mapstore"
	r "github.com/rashmi43/go-messenger/router"
	"log"
	"net/http"
	"time"
)

func main() {
	controller := &cc.MessageController{ //Facade
		Store: mapstore.NewMapStore(), // Inject the dependency
		// store: = mongodb.NewMongoStore()
	}

	router := r.SetMessageRoutes(controller)
	// A Server defines parameters for running an HTTP server.
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port 8001")
	// ListenAndServe listens on the TCP network address and
	// then calls Serve to handle requests on incoming connections.
	log.Fatal(srv.ListenAndServe())

}
