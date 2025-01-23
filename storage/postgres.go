package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Driver de PostgreSQL
	"log"
	"url-shortener/utils"
)

var db *sql.DB

func init() {
	// Conecta a PostgreSQL
	var err error
	connStr := "user=j0k3r dbname=urlshortener sslmode=disable password=1234"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a PostgreSQL")

	// Crea la tabla si no existe
	createTableIfNotExists()
}

// createTableIfNotExists crea la tabla short_urls si no existe
func createTableIfNotExists() {
	query := `
	CREATE TABLE IF NOT EXISTS short_urls (
		short_id VARCHAR(10) PRIMARY KEY,
		is_used BOOLEAN DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creando la tabla short_urls: %v", err)
	}
	fmt.Println("Tabla short_urls creada o ya existente")
}

// GetRandomAvailableKey obtiene una clave aleatoria disponible
func GetRandomAvailableKey() (string, error) {
	for {
		// Genera una clave aleatoria
		key := utils.GenerateRandomKey()

		// Verifica si la clave ya está en uso
		var isUsed bool
		err := db.QueryRow("SELECT is_used FROM short_urls WHERE short_id = $1", key).Scan(&isUsed)
		if err == sql.ErrNoRows {
			// La clave no existe, así que está disponible
			_, err = db.Exec("INSERT INTO short_urls (short_id, is_used) VALUES ($1, $2)", key, true)
			if err != nil {
				return "", fmt.Errorf("error insertando la clave: %v", err)
			}
			return key, nil
		} else if err != nil {
			return "", fmt.Errorf("error verificando la clave: %v", err)
		}

		// Si la clave ya está en uso, intenta con otra
	}
}