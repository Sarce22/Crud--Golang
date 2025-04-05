package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DB es una variable global que almacena la referencia a la base de datos MongoDB.
var DB *mongo.Database

// ConnectDB establece la conexión con la base de datos MongoDB.
// Lee las variables de entorno `MONGO_URI` y `MONGO_DATABASE` para configurar la conexión.
// Si las variables no están definidas o ocurre un error, el programa termina con un log.Fatal.
func ConnectDB() {
    // Leer variables del entorno para obtener la URI de MongoDB y el nombre de la base de datos.
    mongoURI := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DATABASE")

    // Validar que las variables de entorno estén definidas.
    if mongoURI == "" || dbName == "" {
        log.Fatal("❌ Error: Variables de entorno MONGO_URI o MONGO_DATABASE no están definidas")
    }

    // Configurar las opciones del cliente de MongoDB utilizando la URI.
    clientOptions := options.Client().ApplyURI(mongoURI)

    // Crear un nuevo cliente de MongoDB.
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err) // Terminar el programa si ocurre un error al crear el cliente.
    }

    // Crear un contexto con un tiempo límite de 10 segundos para la conexión.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() // Asegurarse de liberar el contexto al finalizar.

    // Conectar el cliente al servidor de MongoDB.
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err) // Terminar el programa si ocurre un error al conectar.
    }

    // Asignar la base de datos especificada a la variable global DB.
    DB = client.Database(dbName)
    fmt.Println("✅ Conectado a MongoDB:", dbName) // Confirmar la conexión exitosa.
}