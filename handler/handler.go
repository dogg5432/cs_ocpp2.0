package handler

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

type CSMSHandler struct {
	ChargingStations map[string]*ChargingStationState
}

type ConnectorInfo struct {
	status             availability.ConnectorStatus
	currentTransaction int
}

type ChargingStationState struct {
	Status         availability.ChangeAvailabilityStatus
	FirmwareStatus firmware.FirmwareStatus
	Connectors     map[int]*ConnectorInfo
	Transactions   map[int]*TransactionInfo
}

type TransactionInfo struct {
	id          int
	startTime   *types.DateTime
	endTime     *types.DateTime
	startMeter  int
	endMeter    int
	connectorID int
	idTag       string
}