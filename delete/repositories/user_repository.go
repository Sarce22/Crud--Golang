package repositories

import (
    "context"
    "crud-microservice/config"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// UserRepository proporciona métodos para interactuar con la colección de usuarios en MongoDB.
type UserRepository struct {
    Collection *mongo.Collection // Referencia a la colección "users" en la base de datos.
}

// NewUserRepository crea una nueva instancia de UserRepository.
// Inicializa la colección "users" desde la configuración de la base de datos.
func NewUserRepository() *UserRepository {
    return &UserRepository{
        Collection: config.DB.Collection("users"),
    }
}

// DeleteUserByID elimina un usuario por su ID.
// Recibe como parámetro un `primitive.ObjectID` que representa el ID del usuario.
// Retorna un error si ocurre un problema o si no se encuentra el usuario.
func (r *UserRepository) DeleteUserByID(id primitive.ObjectID) error {
    // Crear un contexto con un tiempo límite de 5 segundos.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Filtro para buscar el usuario por ID.
    filter := bson.M{"_id": id}

	// Intentar eliminar el usuario de la colección.
    result, err := r.Collection.DeleteOne(ctx, filter)
    if err != nil {
        // Retornar el error si ocurre un problema durante la eliminación.
        return err
    }

    // Verificar si no se eliminó ningún documento.
    if result.DeletedCount == 0 {
        return mongo.ErrNoDocuments // Retornar un error indicando que no se encontró el usuario.
    }

    // Retornar nil si la eliminación fue exitosa.
    return nil
}

// DeleteUserByCedula elimina un usuario por su cédula.
// Recibe como parámetro un string que representa la cédula del usuario.
// Retorna un error si ocurre un problema o si no se encuentra el usuario.
func (r *UserRepository) DeleteUserByCedula(cedula string) error {
    // Crear un contexto con un tiempo límite de 5 segundos.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Filtro para buscar el usuario por cédula.
    filter := bson.M{"cedula": cedula}

    // Intentar eliminar el usuario de la colección.
    result, err := r.Collection.DeleteOne(ctx, filter)
    if err != nil {
        // Retornar el error si ocurre un problema durante la eliminación.
        return err
    }

    // Verificar si no se eliminó ningún documento.
    if result.DeletedCount == 0 {
        return mongo.ErrNoDocuments // Retornar un error indicando que no se encontró el usuario.
    }

    // Retornar nil si la eliminación fue exitosa.
    return nil
}