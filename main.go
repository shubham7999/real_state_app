package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"real_state/handlers"
	"real_state/middleware"
	"real_state/utils"
)

func main() {

	utils.InitDatabase()

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/api/v1/houses", handlers.UploadHouse).Methods("POST")
	router.HandleFunc("/api/v1/houses/{id}", handlers.GetHouse).Methods("GET")
	router.HandleFunc("/api/v1/houses", handlers.ListHouses).Methods("GET")
	router.HandleFunc("/api/v1/houses/{id}", handlers.UpdateHouse).Methods("PUT")
	router.HandleFunc("/api/v1/houses/{id}", handlers.DeleteHouse).Methods("DELETE")


	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
