package controllers

import (
	"encoding/json"
	"net/http"

	"crud-microservice/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserController maneja las solicitudes relacionadas con la eliminación de usuarios.
type UserController struct {
	Service services.IUserServiceInterface // Servicio que contiene la lógica de negocio para usuarios.
}

// NewUserController crea una nueva instancia de UserController.
// Recibe como parámetro un puntero a UserService y lo asocia al controlador.
func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// DeleteUserByID elimina un usuario por su ID.
// Recibe el ID del usuario como parámetro en la URL.
// Responde con un mensaje de éxito o un error si el usuario no existe.
func (c *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de la URL.
	params := mux.Vars(r)
	id := params["id"]

	// Convertir el ID de string a ObjectID.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Responder con un error si el ID es inválido.
		http.Error(w, "❌ ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para eliminar el usuario por ID.
	err = c.Service.DeleteUserByID(objID)
	if err != nil {
		// Responder con un error si el usuario no se encuentra.
		http.Error(w, "❌ No se encontró el usuario con ese ID", http.StatusNotFound)
		return
	}

	// Responder con un mensaje de éxito si el usuario fue eliminado.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Usuario eliminado correctamente"})
}

// DeleteUserByCedula elimina un usuario por su cédula.
// Recibe la cédula del usuario como parámetro en la URL.
// Responde con un mensaje de éxito o un error si el usuario no existe.
func (c *UserController) DeleteUserByCedula(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de la URL.
	params := mux.Vars(r)
	cedula := params["cedula"]

	// Llamar al servicio para eliminar el usuario por cédula.
	err := c.Service.DeleteUserByCedula(cedula)
	if err != nil {
		// Responder con un error si el usuario no se encuentra.
		http.Error(w, "❌ No se encontró el usuario con esa cédula", http.StatusNotFound)
		return
	}

	// Responder con un mensaje de éxito si el usuario fue eliminado.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Usuario eliminado correctamente"})
}
