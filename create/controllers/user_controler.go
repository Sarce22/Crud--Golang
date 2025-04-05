package controllers

import (
	"encoding/json"
	"net/http"

	"crud-microservice/services"
	
)

// UserController es un controlador que maneja las solicitudes relacionadas con usuarios.
type UserController struct {
	Service services.UserServiceInterface // Servicio que contiene la lógica de negocio para usuarios.
}

// NewUserController crea una nueva instancia de UserController.
// Recibe como parámetro un puntero a UserService y lo asocia al controlador.
func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// Response define la estructura estándar para las respuestas HTTP.
// Incluye un indicador de éxito, un mensaje y datos opcionales.
type Response struct {
	Success bool        `json:"success"`        // Indica si la operación fue exitosa.
	Message string      `json:"message"`        // Mensaje descriptivo del resultado.
	Data    interface{} `json:"data,omitempty"` // Datos adicionales (opcional).
}

// CreateUser maneja la creación de un nuevo usuario.
// Recibe una solicitud HTTP con datos en formato JSON y responde con el resultado.
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Estructura para mapear los datos de la solicitud JSON.
	var request struct {
		Nombre    string `json:"nombre"`    // Nombre del usuario.
		Telefono  string `json:"telefono"`  // Teléfono del usuario.
		Direccion string `json:"direccion"` // Dirección del usuario.
		Cedula    string `json:"cedula"`    // Cédula del usuario.
		Correo    string `json:"correo"`    // Correo electrónico del usuario.
	}

	// Decodificar el cuerpo de la solicitud JSON en la estructura request.
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Si ocurre un error al decodificar, responder con un código 400 (Bad Request).
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "❌ Error: El formato del JSON es incorrecto.",
		})
		return
	}

	// Crear usuario
	user, err := c.Service.CreateUser(request.Nombre, request.Telefono, request.Direccion, request.Cedula, request.Correo)
	if err != nil {
		// Si ocurre un error (por ejemplo, usuario duplicado), responder con un código 409 (Conflict).
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Si el usuario se crea exitosamente, responder con un código 201 (Created).
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "✅ Usuario creado exitosamente.",
		Data:    user, // Incluir los datos del usuario creado en la respuesta.
	})
}
