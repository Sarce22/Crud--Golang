// controllers/mocks/user_service_mock.go
package mocks

import (
	"crud-microservice/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(nombre, telefono, direccion, cedula, correo string) (*models.User, error) {
	args := m.Called(nombre, telefono, direccion, cedula, correo)
	return args.Get(0).(*models.User), args.Error(1)
}
