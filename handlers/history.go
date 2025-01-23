package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/storage"
)

func GetUserHistory(w http.ResponseWriter, r *http.Request) {
	// Obtiene el userId del usuario
	userID := SetUserIDCookie(w, r)

	// Obtiene las URLs acortadas por el usuario
	urls, err := storage.GetUserURLs(userID)
	if err != nil {
		http.Error(w, "Error obteniendo el historial", http.StatusInternalServerError)
		return
	}

	// Devuelve la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}