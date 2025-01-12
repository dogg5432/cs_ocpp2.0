package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChargePoint struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty"`
	ChargePointID string                 `bson:"chargePointId"`
	Vendor        string                 `bson:"vendor"`
	Model         string                 `bson:"model"`
	Status        core.ChargePointStatus `bson:"status"`
	Connectors    []Connector            `bson:"connectors"`
	CreatedAt     time.Time              `bson:"createdAt"`
	UpdatedAt     time.Time              `bson:"updatedAt"`
}