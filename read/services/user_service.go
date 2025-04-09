package services

import (
    "crud-microservice/models"
    "crud-microservice/repositories"
)

// UserService proporciona la lógica de negocio para las operaciones relacionadas con la lectura de usuarios.
type UserService struct {
    Repo *repositories.UserRepository // Repositorio para interactuar con la base de datos de usuarios.
}

// IUserServiceInterface define la interfaz para los servicios de usuario.
type IUserServiceInterface interface {
    GetUserByCedula(cedula string) (*models.User, error)
    GetAllUsers() ([]models.User, error)
}

// NewUserService crea una nueva instancia de UserService.
// Recibe como parámetro un repositorio de usuarios y lo asocia al servicio.
func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

// GetAllUsers obtiene todos los usuarios del sistema.
// Llama al repositorio para obtener la lista de usuarios.
// Retorna una lista de usuarios o un error si ocurre un problema.
func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAllUsers()
}

// GetUserByCedula obtiene un usuario por su cédula.
// Recibe como parámetro un string que representa la cédula del usuario.
// Llama al repositorio para buscar el usuario por cédula.
// Retorna el usuario encontrado o un error si no se encuentra.
func (s *UserService) GetUserByCedula(cedula string) (*models.User, error) {
    return s.Repo.GetUserByCedula(cedula)
}

//tets 3