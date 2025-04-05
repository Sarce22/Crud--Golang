package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User representa la estructura de un usuario en el sistema.
// Este modelo se utiliza para mapear los datos almacenados en la base de datos MongoDB.
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`        // Identificador único del usuario (generado automáticamente por MongoDB).
    Nombre    string             `bson:"nombre" json:"nombre"`          // Nombre del usuario.
    Telefono  string             `bson:"telefono" json:"telefono"`      // Número de teléfono del usuario.
    Direccion string             `bson:"direccion" json:"direccion"`    // Dirección física del usuario.
    Cedula    string             `bson:"cedula" json:"cedula"`          // Número de cédula o identificación del usuario.
    Correo    string             `bson:"correo" json:"correo"`          // Correo electrónico del usuario.
}