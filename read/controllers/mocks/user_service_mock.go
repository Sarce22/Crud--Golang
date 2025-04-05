package mocks

import (
    "crud-microservice/models"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUserByCedula(cedula string) (*models.User, error) {
    args := m.Called(cedula)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
    args := m.Called()
    return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) DeleteUserByID(id primitive.ObjectID) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *MockUserService) DeleteUserByCedula(cedula string) error {
    args := m.Called(cedula)
    return args.Error(0)
}
