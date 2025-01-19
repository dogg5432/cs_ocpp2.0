package model

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Charger struct {
	ID              primitive.ObjectID                    `bson:"_id,omitempty"`
	ChargeStationID string                                `bson:"chargeStationID"`
	VendorName      string                                `bson:"vendorName"`
	Model           string                                `bson:"model"`
	Status          availability.ChangeAvailabilityStatus `bson:"status"`
	FirmwareVersion string                                `bson:"firmwareVersion"`
	SerialNumber    string                                `bson:"serialNumber"`
	CreatedAt       time.Time                             `bson:"createdAt"`
	UpdatedAt       time.Time                             `bson:"updatedAt"`
}