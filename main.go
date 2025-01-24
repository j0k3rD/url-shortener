package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/handlers"
	"github.com/rs/cors"
)

func main() {
	// Carga la configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error cargando configuración:", err)
	}

	// Configura CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.Server.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Origin"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Configura los manejadores
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten-url", handlers.ShortenURL)
	mux.HandleFunc("/history", handlers.GetUserHistory)
	mux.HandleFunc("/", handlers.RedirectURL)

	// Aplica el middleware CORS
	handler := c.Handler(mux)

	// Inicia el servidor
	fmt.Printf("Servidor escuchando en el puerto %s...\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, handler))
}