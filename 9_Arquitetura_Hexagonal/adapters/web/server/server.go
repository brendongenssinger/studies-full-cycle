package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codeedu/go-hexagonal/adapters/web/handler"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	_ "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w *WebServer) Server() {

	r := mux.NewRouter()
	// negroni vai trabalhar como middleware
	// log
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is running")
}
