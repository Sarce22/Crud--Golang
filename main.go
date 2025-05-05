package main

import (
	"log"
	"net/http"
)

func main() {
	// Configuración de la ruta
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World and Docker! 🚀"))
	})

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// En este código, hemos creado un servidor HTTP simple que responde con "Hello, World! 🚀" en la ruta raíz.
