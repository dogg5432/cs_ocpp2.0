package repository

import (
	"context"

	"github.com/dogg5432/cs_ocpp2.0/database"
	"github.com/dogg5432/cs_ocpp2.0/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type connectorsRepository struct {
	collection *mongo.Collection
}

func NewConnectorsRepository() ConnectorRepository {
	return &connectorsRepository{
		collection: database.Client.Collection("connectors"),
	}
}

func (c connectorsRepository) Create(connector *model.Connector) error {
	var _, err = c.collection.InsertOne(context.TODO(), connector)
	if err != nil {
		return err
	}
	return nil
}

func (c connectorsRepository) FindOne(connectorModel *model.Connector) (model.Connector, error) {
	var connector model.Connector
	var filter = bson.M{"connectorId": connectorModel.ConnectorID, "chargerId": connectorModel.ChargerID}
	err := c.collection.FindOne(context.TODO(), filter).Decode(&connector)
	if err == mongo.ErrNoDocuments {
		return model.Connector{}, nil
	} else if err != nil {
		panic(err)
	}
	return connector, nil
}

func (c connectorsRepository) Update(connectorModel *model.Connector) error {
	var filter = bson.M{"connectorId": connectorModel.ConnectorID, "chargerId": connectorModel.ChargerID}
	var _, err = c.collection.UpdateOne(context.TODO(), filter, connectorModel)
	if err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		panic(err)
	}
	return nil
}

func (c connectorsRepository) InsertOrUpdate(connector *model.Connector) error {
	checkConnector, err := c.FindOne(connector)
	if err != nil {
		return err
	} else if (checkConnector == model.Connector{}) {
		return c.Create(connector)
	}
	return c.Update(connector)
}
