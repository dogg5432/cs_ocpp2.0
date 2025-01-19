package repository

import "github.com/dogg5432/cs_ocpp2.0/model"

type ChargerRepository interface {
	Create(charger *model.Charger) error
	FindOne(chargerID string) (model.Charger, error)
}

type ConnectorRepository interface {
	Create(connector *model.Connector) error
	FindOne(connectorModel *model.Connector) (model.Connector, error)
	Update(connector *model.Connector) error
	InsertOrUpdate(connector *model.Connector) error
}