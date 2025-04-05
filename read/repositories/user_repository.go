package repositories

import (
    "context"
    "crud-microservice/config"
    "crud-microservice/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
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

// GetAllUsers obtiene todos los usuarios de la base de datos.
// Retorna una lista de usuarios o un error si ocurre un problema.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    // Crear un contexto con un tiempo límite de 5 segundos.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var users []models.User // Lista para almacenar los usuarios obtenidos.

    // Realizar la consulta para obtener todos los documentos de la colección.
    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        // Retornar un error si ocurre un problema durante la consulta.
        return nil, err
    }
    defer cursor.Close(ctx) // Cerrar el cursor al finalizar.

    // Iterar sobre los documentos obtenidos.
    for cursor.Next(ctx) {
        var user models.User
        // Decodificar cada documento en un objeto `User`.
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        // Agregar el usuario a la lista.
        users = append(users, user)
    }

    // Retornar la lista de usuarios.
    return users, nil
}

// GetUserByCedula obtiene un usuario por su cédula.
// Recibe como parámetro un string que representa la cédula del usuario.
// Retorna el usuario encontrado o un error si no se encuentra.
func (r *UserRepository) GetUserByCedula(cedula string) (*models.User, error) {
    // Crear un contexto con un tiempo límite de 5 segundos.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User // Objeto para almacenar el usuario encontrado.
    filter := bson.M{"cedula": cedula} // Filtro para buscar el usuario por cédula.

    // Realizar la consulta para encontrar el usuario.
    err := r.Collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        // Retornar un error si no se encuentra el usuario.
        return nil, err
    }

    // Retornar el usuario encontrado.
    return &user, nil
}