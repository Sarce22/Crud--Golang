package services

import (
    "crud-microservice/repositories"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService proporciona la lógica de negocio para las operaciones relacionadas con la eliminación de usuarios.
type UserService struct {
    Repo *repositories.UserRepository // Repositorio para interactuar con la base de datos de usuarios.
}

type IUserServiceInterface interface {
	DeleteUserByID(id primitive.ObjectID) error
	DeleteUserByCedula(cedula string) error
}

// NewUserService crea una nueva instancia de UserService.
// Recibe como parámetro un repositorio de usuarios y lo asocia al servicio.
func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

// DeleteUserByID elimina un usuario por su ID.
// Recibe como parámetro un `primitive.ObjectID` que representa el ID del usuario.
// Llama al repositorio para realizar la operación de eliminación.
// Retorna un error si ocurre un problema o si el usuario no se encuentra.
func (s *UserService) DeleteUserByID(id primitive.ObjectID) error {
    return s.Repo.DeleteUserByID(id)
}

// DeleteUserByCedula elimina un usuario por su cédula.
// Recibe como parámetro un string que representa la cédula del usuario.
// Llama al repositorio para realizar la operación de eliminación.
// Retorna un error si ocurre un problema o si el usuario no se encuentra.
func (s *UserService) DeleteUserByCedula(cedula string) error {
    return s.Repo.DeleteUserByCedula(cedula)
}

// test 3

//test profe