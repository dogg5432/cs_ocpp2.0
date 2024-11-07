package handler

import (
	"github.com/dogg5432/cs_ocpp2.0/utils"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

var log = utils.Log

// func init() {
// 	log = logrus.New()
// 	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
// 	log.SetLevel(logrus.InfoLevel)
// }

// TransactionInfo contains info about a transaction
type TransactionInfo struct {
	id          int
	startTime   *types.DateTime
	endTime     *types.DateTime
	startMeter  int
	endMeter    int
	connectorID int
	idTag       string
}

func (ti *TransactionInfo) hasTransactionEnded() bool {
	return ti.endTime != nil && !ti.endTime.IsZero()
}

// ConnectorInfo contains status and ongoing transaction ID for a connector
type ConnectorInfo struct {
	status             availability.ConnectorStatus
	currentTransaction int
}

func (ci *ConnectorInfo) hasTransactionInProgress() bool {
	return ci.currentTransaction >= 0
}

// ChargingStationState contains some simple state for a connected charging station
type ChargingStationState struct {
	Status         availability.ChangeAvailabilityStatus
	FirmwareStatus firmware.FirmwareStatus
	Connectors     map[int]*ConnectorInfo
	Transactions   map[int]*TransactionInfo
}

func (s *ChargingStationState) getConnector(id int) *ConnectorInfo {
	ci, ok := s.Connectors[id]
	if !ok {
		ci = &ConnectorInfo{currentTransaction: -1}
		s.Connectors[id] = ci
	}
	return ci
}

// CSMSHandler contains some simple state that a CSMS may want to keep.
// In production this will typically be replaced by database/API calls.
type CSMSHandler struct {
	ChargingStations map[string]*ChargingStationState
}

// Utility functions

