package repositories

import (
    "context"
    "crud-microservice/models"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// UserRepositoryInterface define métodos para el repositorio (permite mocking)
type UserRepositoryInterface interface {
    Exists(cedula, correo string) (bool, error)
    CreateUser(user models.User) (*mongo.InsertOneResult, error)
}

// UserRepository implementa UserRepositoryInterface
type UserRepository struct {
    Collection *mongo.Collection
}

// NewUserRepository crea una nueva instancia de UserRepository.
func NewUserRepository(collection *mongo.Collection) *UserRepository {
    return &UserRepository{Collection: collection}
}

// Exists verifica si ya existe un usuario con la misma cédula o correo.
func (r *UserRepository) Exists(cedula, correo string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"$or": []bson.M{
        {"cedula": cedula},
        {"correo": correo},
    }}

    var existingUser models.User
    err := r.Collection.FindOne(ctx, filter).Decode(&existingUser)
    if err == mongo.ErrNoDocuments {
        return false, nil
    } else if err != nil {
        return false, err
    }

    return true, nil
}

// CreateUser crea un usuario si no hay duplicados.
func (r *UserRepository) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
    exists, err := r.Exists(user.Cedula, user.Correo)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("usuario con la misma cédula o correo ya existe")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    return r.Collection.InsertOne(ctx, user)
}
