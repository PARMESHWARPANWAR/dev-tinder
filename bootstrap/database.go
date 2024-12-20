package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PARMESHWARPANWAR/dev-tinder/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	fmt.Println("DB:=>",dbUser,dbPass)
	// mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.2j1hm.mongodb.net/go-project?retryWrites=true&w=majority", dbUser, dbPass)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}