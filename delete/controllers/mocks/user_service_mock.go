// controllers/mocks/user_service_mock.go

package mocks

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) DeleteUserByID(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) DeleteUserByCedula(cedula string) error {
	args := m.Called(cedula)
	return args.Error(0)
}
