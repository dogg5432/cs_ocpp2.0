package model

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connector struct {
	ID          primitive.ObjectID           `bson:"_id,omitempty"`
	ChargerID   string                       `json:"chargerId"`
	ConnectorID int                          `json:"connectorId"`
	Status      availability.ConnectorStatus `json:"status"`
	CreatedAt   time.Time                    `json:"createdAt"`
	UpdatedAt   time.Time                    `json:"updatedAt"`
}
