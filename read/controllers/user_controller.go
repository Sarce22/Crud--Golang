package controllers

import (
    "encoding/json"
    "net/http"

    "crud-microservice/services"
    "github.com/gorilla/mux"
)

// UserController maneja las solicitudes relacionadas con la lectura de usuarios.
type UserController struct {
    Service services.IUserServiceInterface // Servicio que contiene la lógica de negocio para usuarios.
}

// NewUserController crea una nueva instancia de UserController.
// Recibe como parámetro un puntero a UserService y lo asocia al controlador.
func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

// GetAllUsers obtiene todos los usuarios del sistema.
// Responde con una lista de usuarios en formato JSON o un error si ocurre un problema.
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    // Llamar al servicio para obtener todos los usuarios.
    users, err := c.Service.GetAllUsers()
    if err != nil {
        // Responder con un error si ocurre un problema al obtener los usuarios.
        http.Error(w, "❌ Error al obtener los usuarios", http.StatusInternalServerError)
        return
    }

    // Configurar el encabezado de la respuesta como JSON y enviar la lista de usuarios.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// GetUserByCedula obtiene un usuario por su cédula.
// Recibe la cédula como parámetro en la URL.
// Responde con los datos del usuario en formato JSON o un error si no se encuentra.
func (c *UserController) GetUserByCedula(w http.ResponseWriter, r *http.Request) {
    // Obtener la cédula del usuario desde los parámetros de la URL.
    vars := mux.Vars(r)
    cedula := vars["cedula"]

    // Llamar al servicio para obtener el usuario por cédula.
    user, err := c.Service.GetUserByCedula(cedula)
    if err != nil {
        // Responder con un error si no se encuentra el usuario.
        http.Error(w, "❌ No se encontró ningún usuario con esa cédula", http.StatusNotFound)
        return
    }

    // Configurar el encabezado de la respuesta como JSON y enviar los datos del usuario.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}