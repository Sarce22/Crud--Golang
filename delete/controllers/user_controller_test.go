package controllers

import (
	"crud-microservice/controllers/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUserByID_Success(t *testing.T) {
	mockService := new(mocks.MockUserService)
	controller := UserController{Service: mockService}

	id := primitive.NewObjectID()
	mockService.On("DeleteUserByID", id).Return(nil)

	req := httptest.NewRequest("DELETE", "/usuarios/"+id.Hex(), nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/usuarios/{id}", controller.DeleteUserByID)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUserByID_NotFound(t *testing.T) {
	mockService := new(mocks.MockUserService)
	controller := UserController{Service: mockService}

	id := primitive.NewObjectID()
	mockService.On("DeleteUserByID", id).Return(errors.New("usuario no encontrado"))

	req := httptest.NewRequest("DELETE", "/usuarios/"+id.Hex(), nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/usuarios/{id}", controller.DeleteUserByID)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUserByCedula_Success(t *testing.T) {
	mockService := new(mocks.MockUserService)
	controller := UserController{Service: mockService}

	cedula := "12345678"
	mockService.On("DeleteUserByCedula", cedula).Return(nil)

	req := httptest.NewRequest("DELETE", "/usuarios/cedula/"+cedula, nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/usuarios/cedula/{cedula}", controller.DeleteUserByCedula)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUserByCedula_NotFound(t *testing.T) {
	mockService := new(mocks.MockUserService)
	controller := UserController{Service: mockService}

	cedula := "00000000"
	mockService.On("DeleteUserByCedula", cedula).Return(errors.New("usuario no encontrado"))

	req := httptest.NewRequest("DELETE", "/usuarios/cedula/"+cedula, nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/usuarios/cedula/{cedula}", controller.DeleteUserByCedula)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	mockService.AssertExpectations(t)
}
