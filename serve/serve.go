package serve

import (
	"github.com/dogg5432/cs_ocpp2.0/handler"
	"github.com/dogg5432/cs_ocpp2.0/config"
	"github.com/dogg5432/cs_ocpp2.0/util"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1"
	"github.com/lorenzodonini/ocpp-go/ocppj"
)

var log  = util.Log

func Run() {
	configApp := config.ConfigApp.Server
	csms := ocpp2.NewCSMS(nil, nil)
	// Set callback handlers for connect/disconnect
	csms.SetNewChargingStationHandler(func(chargingStation ocpp2.ChargingStationConnection) {
		log.Printf("new charging station %v connected", chargingStation.ID())
	})
	csms.SetChargingStationDisconnectedHandler(func(chargingStation ocpp2.ChargingStationConnection) {
		log.Printf("charging station %v disconnected", chargingStation.ID())
	})
	// Set handler for profile callbacks
	ocppHandlers := &handler.CSMSHandler{ChargingStations: map[string]*handler.ChargingStationState{}}
	csms.SetAuthorizationHandler(ocppHandlers)
	csms.SetAvailabilityHandler(ocppHandlers)
	csms.SetDiagnosticsHandler(ocppHandlers)
	csms.SetFirmwareHandler(ocppHandlers)
	csms.SetLocalAuthListHandler(ocppHandlers)
	csms.SetMeterHandler(ocppHandlers)
	csms.SetProvisioningHandler(ocppHandlers)
	csms.SetRemoteControlHandler(ocppHandlers)
	csms.SetReservationHandler(ocppHandlers)
	csms.SetTariffCostHandler(ocppHandlers)
	csms.SetTransactionsHandler(ocppHandlers)
	// Add handlers for dis/connection of charging stations
	csms.SetNewChargingStationHandler(func(chargingStation ocpp2.ChargingStationConnection) {
		ocppHandlers.ChargingStations[chargingStation.ID()] = &handler.ChargingStationState{Connectors: map[int]*handler.ConnectorInfo{}, Transactions: map[int]*handler.TransactionInfo{}}
		log.WithField("client", chargingStation.ID()).Info("new charging station connected")
		// go exampleRoutine(chargingStation.ID(), handler)
	})
	csms.SetChargingStationDisconnectedHandler(func(chargingStation ocpp2.ChargingStationConnection) {
		log.WithField("client", chargingStation.ID()).Info("charging station disconnected")
		delete(ocppHandlers.ChargingStations, chargingStation.ID())
	})
	ocppj.SetLogger(log)
	// Run CSMS
	log.Infof("starting CSMS on port %d", configApp.Port)
	csms.Start(configApp.Port, "/ocpp2/{chargepoint}")
	log.Info("stopped CSMS")
}
