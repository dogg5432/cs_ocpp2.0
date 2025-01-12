package repository

import (
	"context"

	"github.com/dogg5432/cs_ocpp2.0/database"
	"github.com/dogg5432/cs_ocpp2.0/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type chargerRepository struct {
	collection *mongo.Collection
}

func NewChagersRepository() ChargerRepository {
	return &chargerRepository{
		collection: database.Client.Collection("chargers"),
	}
}

func (c chargerRepository) Create(charger *model.Charger) error {
	var _, err = c.collection.InsertOne(context.TODO(), charger)
	if err != nil {
		return err
	}
	return nil
}
