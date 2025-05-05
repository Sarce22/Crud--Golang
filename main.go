package main

import (
	"log"
	"net/http"
)

func main() {
	// Configuración de la ruta raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World and Docker! 🚀"))
	})

	// Configuración de la ruta /hola
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Hola Mundo! 🌎"))
	})

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


//test