package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/liberopassadorneto/auction/configuration/database/mongodb"
	"log"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal(err)
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
