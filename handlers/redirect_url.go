package handlers

import (
	"net/http"
	"url-shortener/storage"
	"time"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// Extrae el ID corto de la URL
	shortID := r.URL.Path[1:] // Elimina el "/" inicial

	// Busca la URL larga en MongoDB
	longURL, err := storage.GetURL(shortID)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Incrementa el contador de clicks
	err = storage.IncrementClicks(shortID)
	if err != nil {
		http.Error(w, "Error updating clicks", http.StatusInternalServerError)
		return
	}

	// Recopila metadatos
	metadata := map[string]interface{}{
		"userAgent": r.UserAgent(), // Navegador y sistema operativo del usuario
		"ipAddress": r.RemoteAddr,  // Dirección IP del usuario
		"timestamp": time.Now().Format(time.RFC3339), // Fecha y hora de acceso
	}

	// Actualiza el campo "metadata" en MongoDB
	err = storage.UpdateMetadata(shortID, metadata)
	if err != nil {
		http.Error(w, "Error updating metadata", http.StatusInternalServerError)
		return
	}

	// Redirige a la URL larga
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}