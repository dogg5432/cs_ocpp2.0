package database

import (
	"context"
	"fmt"

	"github.com/dogg5432/cs_ocpp2.0/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Database

func Connect() error {
	clientOptions := options.Client().ApplyURI(config.ConfigApp.Database.Uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if(err != nil){
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	Client = client.Database("central_system")

	return nil
}