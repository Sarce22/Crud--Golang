package controllers

import (
	"bytes"
	"crud-microservice/controllers/mocks"
	"crud-microservice/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser_Success(t *testing.T) {
	mockService := &mocks.MockUserService{}
	controller := UserController{Service: mockService}

	requestBody := map[string]string{
		"nombre":    "Juan",
		"telefono":  "123456789",
		"direccion": "Calle 123",
		"cedula":    "12345678",
		"correo":    "juan@example.com",
	}
	jsonData, _ := json.Marshal(requestBody)

	// Simular el usuario esperado
	expectedUser := &models.User{
		ID:        primitive.NewObjectID(),
		Nombre:    "Juan",
		Telefono:  "123456789",
		Direccion: "Calle 123",
		Cedula:    "12345678",
		Correo:    "juan@example.com",
	}

	mockService.
		On("CreateUser", "Juan", "123456789", "Calle 123", "12345678", "juan@example.com").
		Return(expectedUser, nil)

	req, _ := http.NewRequest("POST", "/usuarios", bytes.NewBuffer(jsonData))
	rr := httptest.NewRecorder()

	controller.CreateUser(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	// Comprobaciones
	assert.True(t, response.Success)
	assert.Equal(t, "✅ Usuario creado exitosamente.", response.Message)

	// Opcional: puedes verificar también que el ID exista
	dataMap := response.Data.(map[string]interface{})
	assert.NotEmpty(t, dataMap["id"])

	mockService.AssertExpectations(t)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	mockService := &mocks.MockUserService{}
	controller := UserController{Service: mockService}

	// JSON inválido (falta comillas y llaves)
	invalidJSON := []byte(`{nombre: Juan,telefono:123`)

	req, _ := http.NewRequest("POST", "/usuarios", bytes.NewBuffer(invalidJSON))
	rr := httptest.NewRecorder()

	controller.CreateUser(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "❌ Error: El formato del JSON es incorrecto.", response.Message)

	// No se debe haber llamado a CreateUser del servicio
	mockService.AssertExpectations(t)
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	mockService := &mocks.MockUserService{}
	controller := UserController{Service: mockService}

	requestBody := map[string]string{
		"nombre":    "Ana",
		"telefono":  "987654321",
		"direccion": "Carrera 45",
		"cedula":    "87654321",
		"correo":    "ana@example.com",
	}
	jsonData, _ := json.Marshal(requestBody)

	// Simulamos un error como si el usuario ya existiera
	mockService.
		On("CreateUser", "Ana", "987654321", "Carrera 45", "87654321", "ana@example.com").
		Return(&models.User{}, errors.New("❌ Error: La cédula o correo ya están registrados."))

	req, _ := http.NewRequest("POST", "/usuarios", bytes.NewBuffer(jsonData))
	rr := httptest.NewRecorder()

	controller.CreateUser(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)

	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "❌ Error: La cédula o correo ya están registrados.", response.Message)

	mockService.AssertExpectations(t)
}
