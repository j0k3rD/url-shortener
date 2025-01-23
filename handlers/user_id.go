package handlers

import (
	"net/http"
	"url-shortener/utils"
)

// SetUserIDCookie establece una cookie con un userId único
func SetUserIDCookie(w http.ResponseWriter, r *http.Request) string {
	// Verifica si el usuario ya tiene una cookie con un userId
	cookie, err := r.Cookie("userId")
	if err == nil {
		// Si ya tiene una cookie, devuelve el userId
		return cookie.Value
	}

	// Si no tiene una cookie, genera un nuevo userId
	userID := utils.GenerateUserID()

	// Crea una nueva cookie con el userId
	cookie = &http.Cookie{
		Name:     "userId",
		Value:    userID,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 año de duración
		HttpOnly: true,               // La cookie no es accesible desde JavaScript
		Secure:   false,              // Cambia a true si usas HTTPS
		SameSite: http.SameSiteLaxMode,
	}

	// Establece la cookie en la respuesta
	http.SetCookie(w, cookie)

	return userID
}