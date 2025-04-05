package repositories

import (
    "context"
    "crud-microservice/models"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// UserRepository proporciona métodos para interactuar con la colección de usuarios en MongoDB.
type UserRepository struct {
    Collection *mongo.Collection // Referencia a la colección "users" en la base de datos.
}

// NewUserRepository crea una nueva instancia de UserRepository.
// Recibe como parámetro una colección de MongoDB y la asocia al repositorio.
func NewUserRepository(collection *mongo.Collection) *UserRepository {
    return &UserRepository{Collection: collection}
}

// UpdateUserByCedula actualiza un usuario por su cédula.
// Recibe como parámetros:
// - `cedula`: La cédula del usuario a actualizar.
// - `updateData`: Un objeto `User` con los datos a actualizar.
// Retorna el resultado de la operación de actualización o un error si ocurre un problema.
func (r *UserRepository) UpdateUserByCedula(cedula string, updateData models.User) (*mongo.UpdateResult, error) {
    // Crear un contexto con un tiempo límite de 5 segundos.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Filtro para buscar el usuario por cédula.
    filter := bson.M{"cedula": cedula}

    // Datos a actualizar en el documento.
    update := bson.M{"$set": bson.M{
        "nombre":    updateData.Nombre,
        "telefono":  updateData.Telefono,
        "direccion": updateData.Direccion,
        "correo":    updateData.Correo,
    }}

    // Logs para depuración.
    log.Println("🔍 Buscando usuario con cédula:", cedula)
    log.Println("📢 Datos a actualizar:", update)

    // Intentar actualizar el usuario en la colección.
    result, err := r.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        // Log del error si ocurre un problema durante la actualización.
        log.Println("❌ Error al actualizar usuario:", err)
    }
    return result, err
}