package controllers

import (
    "encoding/json"
    "log"
    "net/http"

    "crud-microservice/models"
    "crud-microservice/services"

    "github.com/gorilla/mux"
)

// UserController maneja las solicitudes relacionadas con la actualización de usuarios.
type UserController struct {
    Service *services.UserService // Servicio que contiene la lógica de negocio para usuarios.
}

// NewUserController crea una nueva instancia de UserController.
// Recibe como parámetro un puntero a UserService y lo asocia al controlador.
func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

// UpdateUserByCedula actualiza un usuario por su cédula.
// Recibe la cédula como parámetro en la URL y los datos a actualizar en el cuerpo de la solicitud.
// Responde con un mensaje de éxito o un error si ocurre un problema.
func (c *UserController) UpdateUserByCedula(w http.ResponseWriter, r *http.Request) {
    // Obtener la cédula del usuario desde los parámetros de la URL.
    vars := mux.Vars(r)
    cedula := vars["cedula"]

    // Log para verificar la cédula recibida.
    log.Println("📌 Cédula recibida en la URL:", cedula)

    // Decodificar los datos de actualización desde el cuerpo de la solicitud.
    var updateData models.User
    err := json.NewDecoder(r.Body).Decode(&updateData)
    if err != nil {
        // Responder con un error si ocurre un problema al decodificar el JSON.
        http.Error(w, "❌ Error al decodificar JSON", http.StatusBadRequest)
        return
    }

    // Llamar al servicio para actualizar el usuario.
    result, err := c.Service.UpdateUserByCedula(cedula, updateData)
    if err != nil {
        // Responder con un error si ocurre un problema durante la actualización.
        http.Error(w, "❌ Error al actualizar usuario", http.StatusInternalServerError)
        return
    }

    // Si no se encontró ningún usuario con la cédula proporcionada.
    if result.MatchedCount == 0 {
        http.Error(w, "❌ No se encontró ningún usuario con esa cédula", http.StatusNotFound)
        return
    }

    // Responder con un mensaje de éxito si el usuario fue actualizado correctamente.
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "✅ Usuario actualizado correctamente"})
}