package main

import (
	"log"
	"net/http"

	"crud-microservice/config"
	"crud-microservice/controllers"
	"crud-microservice/repositories"
	"crud-microservice/services"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	router := mux.NewRouter()

	
	router.HandleFunc("/users/{id}", controller.DeleteUserByID).Methods("DELETE")
	router.HandleFunc("/users/cedula/{cedula}", controller.DeleteUserByCedula).Methods("DELETE")


	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
