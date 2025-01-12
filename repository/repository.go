package repository

import "github.com/dogg5432/cs_ocpp2.0/model"

type ChargerRepository interface {
	Create(charger *model.Charger) error
}