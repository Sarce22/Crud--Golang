package services

import (
	"crud-microservice/models"
	"crud-microservice/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repository: repo}
}

func (s *UserService) UpdateUserByCedula(cedula string, updateData models.User) (*mongo.UpdateResult, error) {
	return s.Repository.UpdateUserByCedula(cedula, updateData)
}


// test profe