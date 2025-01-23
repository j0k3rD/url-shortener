package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/storage"
	"url-shortener/utils"
	"fmt"	
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	// Obtiene el userId del usuario
	userID := SetUserIDCookie(w, r)

	// Decodifica la solicitud
	var request struct {
		LongURL string `json:"longUrl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Valida la URL
	if !utils.IsValidURL(request.LongURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Obtiene una clave aleatoria disponible
	shortID, err := storage.GetRandomAvailableKey()
	if err != nil {
		http.Error(w, "Error generating short ID", http.StatusInternalServerError)
		return
	}

	// Almacena la URL en MongoDB
	if err := storage.SaveURL(shortID, request.LongURL, userID); err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	// Devuelve la respuesta
	response := struct {
		ShortURL string `json:"shortUrl"`
	}{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortID),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}