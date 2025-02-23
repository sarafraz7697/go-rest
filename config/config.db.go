package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", GetEnvOrFatal("DB_USER"), GetEnvOrFatal("DB_PASS"), GetEnvOrFatal("DB_HOST"), GetEnvOrFatal("DB_PORT"))
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB:", err)
	}

	// Ping database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ Could not ping MongoDB:", err)
	}

	DB = client.Database(GetEnvOrFatal("DB_NAME"))
	fmt.Println("✅ Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Database connection is not initialized. Call ConnectDB() first!")
	}
	return DB.Collection(name)
}
