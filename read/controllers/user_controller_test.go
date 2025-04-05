package controllers

import (
    "crud-microservice/controllers/mocks"
    "crud-microservice/models"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestGetUserByCedula_Success(t *testing.T) {
    mockService := new(mocks.MockUserService)
    controller := UserController{Service: mockService}

    expectedUser := &models.User{
        Nombre:    "Sebastián",
        Telefono:  "3012345678",
        Direccion: "Calle Falsa 123",
        Cedula:    "1234567890",
        Correo:    "sebastian@example.com",
    }

    mockService.On("GetUserByCedula", "1234567890").Return(expectedUser, nil)

    req, _ := http.NewRequest("GET", "/usuarios/cedula/1234567890", nil)
    rr := httptest.NewRecorder()

    router := mux.NewRouter()
    router.HandleFunc("/usuarios/cedula/{cedula}", controller.GetUserByCedula)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var actualUser models.User
    err := json.NewDecoder(rr.Body).Decode(&actualUser)
    assert.NoError(t, err)
    assert.Equal(t, *expectedUser, actualUser)

    mockService.AssertExpectations(t)
}

func TestGetUserByCedula_NotFound(t *testing.T) {
    mockService := new(mocks.MockUserService)
    controller := UserController{Service: mockService}

    mockService.On("GetUserByCedula", "9999999999").Return(&models.User{}, errors.New("usuario no encontrado"))

    req, _ := http.NewRequest("GET", "/usuarios/cedula/9999999999", nil)
    rr := httptest.NewRecorder()

    router := mux.NewRouter()
    router.HandleFunc("/usuarios/cedula/{cedula}", controller.GetUserByCedula)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNotFound, rr.Code)
    mockService.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
    mockService := new(mocks.MockUserService)
    controller := UserController{Service: mockService}

    expectedUsers := []models.User{
        {
            Nombre:    "Ana",
            Telefono:  "3000000000",
            Direccion: "Av. Siempre Viva 742",
            Cedula:    "1111111111",
            Correo:    "ana@example.com",
        },
        {
            Nombre:    "Luis",
            Telefono:  "3111111111",
            Direccion: "Calle Luna Calle Sol",
            Cedula:    "2222222222",
            Correo:    "luis@example.com",
        },
    }

    mockService.On("GetAllUsers").Return(expectedUsers, nil)

    req, _ := http.NewRequest("GET", "/usuarios", nil)
    rr := httptest.NewRecorder()

    router := mux.NewRouter()
    router.HandleFunc("/usuarios", controller.GetAllUsers)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var actualUsers []models.User
    err := json.NewDecoder(rr.Body).Decode(&actualUsers)
    assert.NoError(t, err)
    assert.Equal(t, expectedUsers, actualUsers)

    mockService.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
    mockService := new(mocks.MockUserService)
    controller := UserController{Service: mockService}

    mockService.On("GetAllUsers").Return([]models.User{}, errors.New("fallo en la base de datos"))

    req, _ := http.NewRequest("GET", "/usuarios", nil)
    rr := httptest.NewRecorder()

    router := mux.NewRouter()
    router.HandleFunc("/usuarios", controller.GetAllUsers)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusInternalServerError, rr.Code)
    mockService.AssertExpectations(t)
}
