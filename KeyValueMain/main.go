package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	repository *Repository
	service    *Service
	controller *Controller
)

func main() {
	fmt.Println("=================== Init Controller ===========================")
	fmt.Println("repository init")
	repository = NewRepository("../keyValue.db")
	fmt.Println("service init")
	service = NewService(repository)
	fmt.Println("controller init")
	controller = NewController(service)
	fmt.Println("===============================================================")
	router := mux.NewRouter()
	router.HandleFunc("/api/put", controller.put).Methods("PUT")
	router.HandleFunc("/api/get", controller.get).Methods("GET")
	router.HandleFunc("/api/delete", controller.delete).Methods("DELETE")
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
