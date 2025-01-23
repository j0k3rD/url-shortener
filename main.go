package main

import (
	"log"
	"net/http"
	"fmt"
	"url-shortener/handlers"
)

func main() {
	// Configura los endpoints
	http.HandleFunc("/shorten-url", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	// Inicia el servidor
	fmt.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

	// Sirve los archivos estáticos del frontend
	// fs := http.FileServer(http.Dir("../url-shortener-frontend/dist"))
	// http.Handle("/", fs)