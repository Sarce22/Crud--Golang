package services

import (
	"crud-microservice/models"
	"crud-microservice/repositories"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceInterface interface {
	CreateUser(nombre, telefono, direccion, cedula, correo string) (*models.User, error)
}

// UserService proporciona la lógica de negocio para las operaciones relacionadas con usuarios.
type UserService struct {
	Repo repositories.UserRepositoryInterface // Repositorio para interactuar con la base de datos de usuarios.
}

// NewUserService crea una nueva instancia de UserService.
// Recibe como parámetro un repositorio de usuarios y lo asocia al servicio.
func NewUserService(repo repositories.UserRepositoryInterface) *UserService {
	return &UserService{Repo: repo}
}

// CreateUser crea un nuevo usuario en el sistema.
// Realiza validaciones básicas, verifica duplicados y delega la creación al repositorio.
// Recibe como parámetros los datos del usuario: nombre, teléfono, dirección, cédula y correo.
// Retorna el usuario creado o un error en caso de fallo.
func (s *UserService) CreateUser(nombre, telefono, direccion, cedula, correo string) (*models.User, error) {
	// Validaciones básicas
	if nombre == "" || telefono == "" || direccion == "" || cedula == "" || correo == "" {
		// Retornar un error si algún campo está vacío.
		return nil, errors.New("❌ Todos los campos son obligatorios")
	}

	// Validar si ya existe un usuario con la misma cédula o correo.
	exists, err := s.Repo.Exists(cedula, correo)
	if err != nil {
		// Retornar un error si ocurre un problema al verificar la existencia del usuario.
		return nil, errors.New("❌ Error al verificar la existencia del usuario en la base de datos")
	}
	if exists {
		// Retornar un error si ya existe un usuario con la misma cédula o correo.
		return nil, errors.New("❌ Ya existe un usuario con la misma cédula o correo")
	}

	// Crear un nuevo usuario con los datos proporcionados.
	user := &models.User{
		ID:        primitive.NewObjectID(), // Generar un nuevo ID único para el usuario.
		Nombre:    nombre,
		Telefono:  telefono,
		Direccion: direccion,
		Cedula:    cedula,
		Correo:    correo,
	}

	// Intentar insertar el usuario en la base de datos.
	_, err = s.Repo.CreateUser(*user)
	if err != nil {
		// Retornar un error si ocurre un problema al registrar el usuario.
		return nil, errors.New("❌ Error al registrar el usuario en la base de datos")
	}

	// Retornar el usuario creado si todo fue exitoso.
	return user, nil
}

//prueba de sh

// prueba de sh # 2
