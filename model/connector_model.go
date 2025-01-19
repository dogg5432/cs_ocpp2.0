package model

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
)

type Connector struct {
	ID          string                       `json:"id"`
	ChargerID   string                       `json:"chargerId"`
	ConnectorID int                          `json:"connectorId"`
	Status      availability.ConnectorStatus `json:"status"`
	CreatedAt   time.Time                    `json:"createdAt"`
	UpdatedAt   time.Time                    `json:"updatedAt"`
}
