package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Konfigurasi koneksi MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err.Error())
	}
	defer client.Disconnect(context.Background())

	// Mengambil koneksi ke database MongoDB
	database := client.Database("nama_database")

	// Mengambil koleksi (tabel) dari database
	collection := database.Collection("nama_koleksi")

	// Membaca data dari koleksi
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		panic(err.Error())
	}
	defer cursor.Close(context.Background())

	// Membaca hasil query
	for cursor.Next(context.Background()) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(result)
	}
}
