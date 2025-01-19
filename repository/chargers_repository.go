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

func (c chargerRepository) FindOne(chargerID string) (model.Charger, error) {
	var charger model.Charger
	err := c.collection.FindOne(context.TODO(), model.Charger{ChargeStationID: chargerID}).Decode(&charger)
	if err != nil {
		return model.Charger{}, err
	}
	return charger, nil

}
