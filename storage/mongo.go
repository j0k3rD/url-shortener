package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
	"url-shortener/config"
)

var client *mongo.Client

// init establece la conexión con MongoDB
func init() {
	// Conecta a MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB")
}

func InitMongoDB(cfg *config.Config) error {
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.URI)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	return client.Ping(context.TODO(), nil)
}

// SaveURL guarda una URL en MongoDB con el esquema deseado
func SaveURL(shortID, longURL, userID string) error {
	collection := client.Database("urlshortener").Collection("urls")

	// Crea el documento con el esquema deseado
	document := bson.M{
		"shortUrlId":  shortID,
		"longUrl":     longURL,
		"creationDate": time.Now().Format(time.RFC3339), // Fecha actual en formato ISO 8601
		"userId":      userID, // Asocia la URL con el userId
		"clicks":      0,
		"metadata":    []bson.M{}, // Inicializa metadata como un array vacío
		"isActive":    true,
	}

	// Inserta el documento en la colección
	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return fmt.Errorf("error insertando documento en MongoDB: %v", err)
	}

	return nil
}

// GetURL obtiene la URL larga asociada a un ID corto
func GetURL(shortID string) (string, error) {
	collection := client.Database("urlshortener").Collection("urls")

	// Busca el documento por shortID
	var result struct {
		LongURL string `bson:"longUrl"`
	}
	err := collection.FindOne(context.TODO(), bson.M{"shortUrlId": shortID}).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error obteniendo la URL: %v", err)
	}

	return result.LongURL, nil
}

// GetUserURLs obtiene todas las URLs acortadas por un usuario
func GetUserURLs(userID string) ([]bson.M, error) {
	collection := client.Database("urlshortener").Collection("urls")

	// Busca todos los documentos asociados con el userId
	cursor, err := collection.Find(context.TODO(), bson.M{"userId": userID})
	if err != nil {
		return nil, fmt.Errorf("error obteniendo URLs del usuario: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Decodifica los documentos
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("error decodificando documentos: %v", err)
	}

	return results, nil
}

// IncrementClicks incrementa el contador de clicks de una URL
func IncrementClicks(shortID string) error {
	collection := client.Database("urlshortener").Collection("urls")

	// Incrementa el campo "clicks" en 1
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"shortUrlId": shortID},
		bson.M{"$inc": bson.M{"clicks": 1}},
	)
	if err != nil {
		return fmt.Errorf("error incrementando clicks: %v", err)
	}

	return nil
}

// UpdateMetadata actualiza el campo "metadata" de una URL
func UpdateMetadata(shortID string, metadata map[string]interface{}) error {
	collection := client.Database("urlshortener").Collection("urls")

	// Actualiza el campo "metadata" usando $push
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"shortUrlId": shortID},
		bson.M{"$push": bson.M{"metadata": metadata}},
	)
	if err != nil {
		return fmt.Errorf("error actualizando metadata: %v", err)
	}

	return nil
}